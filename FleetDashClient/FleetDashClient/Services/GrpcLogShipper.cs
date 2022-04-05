using System.Collections.Concurrent;
using EveLogParser;
using EveLogParser.Models.Events;
using EventAggregator.Blazor;
using FleetDashClient.Data;
using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using FleetDashClient.Models.Exceptions;
using FleetDashClient.Protobuf;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Serilog;

namespace FleetDashClient.Services;

public class GrpcLogShipper : IDisposable
{
    private readonly ILogParserService _logParserService;
    private readonly FleetDashService.FleetDashServiceClient _client;
    private readonly IEventAggregator _eventAggregator;
    private readonly IServiceScopeFactory _serviceScopeFactory;
    
    private (string?, DateTimeOffset) _fleetData;
    private readonly ConcurrentDictionary<string, DateTimeOffset> _lastBackOff = new();
    private readonly BlockingCollection<EveLogEvent> _eventQueue = new();

    public GrpcLogShipper(
        ILogParserService logParserService,
        FleetDashService.FleetDashServiceClient client,
        IEventAggregator aggregator,
        IServiceScopeFactory scopeFactory)
    {
        _logParserService = logParserService;
        _client = client;
        _eventAggregator = aggregator;
        _serviceScopeFactory = scopeFactory;

        _logParserService.OnIncomingDamage += EveLogHandler;
        _logParserService.OnOutgoingDamage += EveLogHandler;
        _logParserService.OnIncomingArmor += EveLogHandler;
        _logParserService.OnOutgoingArmor += EveLogHandler;
        _logParserService.OnIncomingShield += EveLogHandler;
        _logParserService.OnOutgoingShield += EveLogHandler;
        _logParserService.OnIncomingHull += EveLogHandler;
        _logParserService.OnOutgoingHull += EveLogHandler;
        _logParserService.OnIncomingCapacitor += EveLogHandler;
        _logParserService.OnOutgoingCapacitor += EveLogHandler;
        _logParserService.OnIncomingNos += EveLogHandler;
        _logParserService.OnOutgoingNos += EveLogHandler;
        _logParserService.OnIncomingNeut += EveLogHandler;
        _logParserService.OnOutgoingNeut += EveLogHandler;
        _logParserService.OnIncomingJam += EveLogHandler;
        _logParserService.OnOutgoingJam += EveLogHandler;

        Task.Run(QueueProcessor);
    }

    private async Task<Token?> GetFreshTokenAsync(DataContext dbContext, string characterId)
    {
        var token = dbContext.Tokens
                .OrderByDescending(x => x.ExpiresAt)
                .FirstOrDefault(x => x.CharacterId == characterId);
        if (token == null)
        {
            return null;
        }

        if (token.ExpiresAt > DateTimeOffset.Now.AddMinutes(1))
        {
            return token;
        }

        // refresh token
        var newToken = await TokenService.RefreshToken(token);
        newToken.CharacterId = characterId;
        dbContext.Tokens.Add(newToken);
        await dbContext.SaveChangesAsync();

        return newToken;
    }

    private void EveLogHandler<T>(object? sender, T args) where T : EveLogEventArgs
    {
        Task.Run(async () =>
        {
            using var scope = _serviceScopeFactory.CreateScope();
            var dbContext = scope.ServiceProvider.GetRequiredService<DataContext>();
            
            var character = await dbContext.Characters.FindAsync(args.CharacterId);
            if (character == null)
            {
                Log.Error("Character not found in DB! ID: {CharacterId}", args.CharacterId);
                return;
            }

            Token? freshToken;
            try
            {
                freshToken = await GetFreshTokenAsync(dbContext, character.Id);
                if (freshToken == null)
                {
                    Log.Warning("Character {CharacterId} has no token", character.Id);
                    await _eventAggregator.PublishAsync(new InvalidCharacterTokenEventArgs(character.Id));
                    return;
                }
            }
            catch (TokenRefreshException ex)
            {
                await _eventAggregator.PublishAsync(new InvalidCharacterTokenEventArgs(character.Id));
                Log.Warning(ex, "Failed to refresh token for character {CharacterId}", character.Id);
                return;
            }

            // Need to be in a fleet to ship logs
            // Eve API caches the route for 60 seconds
            if (_fleetData.Item2 < DateTimeOffset.Now.AddSeconds(-60))
            {
                var fleetId =
                    await EveClient.GetCharacterFleetIdAsync(freshToken.CharacterId, freshToken.AccessToken);
                _fleetData = (fleetId, DateTimeOffset.Now);
            }

            if (_fleetData.Item1 == null)
            {
                Log.Debug("Not in a fleet, skipping log ship");
                return;
            }

            var shipTypeId = await EveClient.GetCharacterShipTypeId(freshToken.CharacterId, freshToken.AccessToken);

            var eventType = args.GetType().ToString().Split('.').Last();
            var eveEvent = new EveLogEvent
            {
                Event = eventType,
                Timestamp = args.Timestamp.ToTimestamp(),
                CharacterId = args.CharacterId,
                CharacterShipTypeId = shipTypeId ?? 0,
                Amount = args.Amount,
                Pilot = args.Pilot,
                Ship = args.Ship,
                Weapon = args.Weapon,
                Application = args.Application,
                Corporation = args.Corporation,
                Alliance = args.Alliance,
            };
            _eventQueue.Add(eveEvent);
            await _eventAggregator.PublishAsync(new LogStreamedEventArgs(args.CharacterId));
        });
    }

    private async Task QueueProcessor()
    {

        using var scope = _serviceScopeFactory.CreateScope();
        var dbContext = scope.ServiceProvider.GetRequiredService<DataContext>();

        while (!_eventQueue.IsCompleted)
        {
            try
            {
                Thread.Sleep(500);
                var batch = new List<EveLogEvent>();
                while (_eventQueue.TryTake(out var eveEvent) && batch.Count < 100)
                {
                    batch.Add(eveEvent);
                }

                if (batch.Count == 0)
                {
                    continue;
                }

                var grouped = batch
                    .GroupBy(x => x.CharacterId)
                    .ToList();

                foreach (var batched in grouped)
                {
                    if (_lastBackOff.TryGetValue(batched.Key, out var lastBackoff) &&
                        lastBackoff > DateTimeOffset.Now.AddSeconds(-30))
                    {
                        return;
                    }

                    try
                    {
                        var freshToken = await GetFreshTokenAsync(dbContext, batched.Key);
                        if (freshToken == null)
                        {
                            Log.Warning("Character {CharacterId} has no token", batched.Key);
                            await _eventAggregator.PublishAsync(new InvalidCharacterTokenEventArgs(batched.Key));
                            return;
                        }
                        
                        var meta = new Metadata
                        {
                            { "Authorization", $"Bearer {freshToken.AccessToken}" }
                        };

                        var batchedRequest = new EveLogEventBatch
                        {
                            Events = { batched }
                        };
                        await _client.PostEveLogEventBatchAsync(batchedRequest, meta);
                    }
                    catch (RpcException ex)
                    {
                        if (ex.StatusCode == StatusCode.FailedPrecondition)
                        {
                            Log.Debug("Server is not accepting logs, backing off for 30 seconds");
                            _lastBackOff.AddOrUpdate(batched.Key, DateTimeOffset.Now, (_, _) => DateTimeOffset.Now);
                        }
                        else
                        {
                            Log.Warning(ex, "Failed to send log event batch to server. Events missed: {BatchCount}",
                                batched.Count());
                        }
                    }
                }
            }
            catch (Exception ex)
            {
                Log.Error(ex, "Failed to process log event queue");
            }
        }
    }
    
    public void Dispose()
    {
        _logParserService.OnIncomingDamage -= EveLogHandler;
        _logParserService.OnOutgoingDamage -= EveLogHandler;
        _logParserService.OnIncomingArmor -= EveLogHandler;
        _logParserService.OnOutgoingArmor -= EveLogHandler;
        _logParserService.OnIncomingShield -= EveLogHandler;
        _logParserService.OnOutgoingShield -= EveLogHandler;
        _logParserService.OnIncomingHull -= EveLogHandler;
        _logParserService.OnOutgoingHull -= EveLogHandler;
        _logParserService.OnIncomingCapacitor -= EveLogHandler;
        _logParserService.OnOutgoingCapacitor -= EveLogHandler;
        _logParserService.OnIncomingNos -= EveLogHandler;
        _logParserService.OnOutgoingNos -= EveLogHandler;
        _logParserService.OnIncomingNeut -= EveLogHandler;
        _logParserService.OnOutgoingNeut -= EveLogHandler;
        _logParserService.OnIncomingJam -= EveLogHandler;
        _logParserService.OnOutgoingJam -= EveLogHandler;
        
        _eventQueue.CompleteAdding();
    }
}