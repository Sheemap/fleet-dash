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
using Microsoft.EntityFrameworkCore;
using Serilog;

namespace FleetDashClient.Services;

public class GrpcLogShipper : IDisposable
{
    private readonly ILogParserService _logParserService;
    private readonly DataContext _dbContext;
    private readonly FleetDashService.FleetDashServiceClient _client;
    private readonly IEventAggregator _eventAggregator;
    
    private (string?, DateTimeOffset) _fleetData;
    private DateTimeOffset _lastBackOff = DateTimeOffset.MinValue;
    private SemaphoreSlim _semaphore = new SemaphoreSlim(1);
    
    public GrpcLogShipper(
        ILogParserService logParserService,
        DataContext dataContext,
        FleetDashService.FleetDashServiceClient client,
        IEventAggregator aggregator)
    {
        _logParserService = logParserService;
        _dbContext = dataContext;
        _client = client;
        _eventAggregator = aggregator;

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
    }

    private async Task<Token?> GetFreshToken(string characterId)
    {
        await _semaphore.WaitAsync();
        
        var token = _dbContext.Tokens
            .OrderByDescending(x => x.ExpiresAt)
            .FirstOrDefault(x => x.CharacterId == characterId);
        if (token == null)
        {
            return null;
        }
     
        if (token.ExpiresAt > DateTimeOffset.Now.AddMinutes(1)) return token;

        // refresh token
        var newToken = await TokenService.RefreshToken(token);
        newToken.CharacterId = characterId;
        _dbContext.Tokens.Add(newToken);
        await _dbContext.SaveChangesAsync();

        _semaphore.Release();
        
        return newToken;
    }

    private void EveLogHandler<T>(object? sender, T args) where T : EveLogEventArgs
    {
        if (_lastBackOff > DateTimeOffset.Now.AddSeconds(-30))
        {
            return;
        }

        Task.Run(async () =>
        {
            var character = await _dbContext.Characters
                .FirstOrDefaultAsync(x => x.Id == args.CharacterId);
            if (character == null)
            {
                Log.Error("Character not found in DB! ID: {CharacterId}", args.CharacterId);
                return;
            }

            Token? freshToken;
            try
            {
                freshToken = await GetFreshToken(character.Id);
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
                var fleetId = await EveClient.GetCharacterFleetIdAsync(freshToken.CharacterId, freshToken.AccessToken);
                _fleetData = (fleetId, DateTimeOffset.Now);
            }
            if (_fleetData.Item1 == null)
            {
                Log.Debug("Not in a fleet, skipping log ship");
                return;
            }

            var meta = new Metadata
            {
                { "Authorization", $"Bearer {freshToken.AccessToken}" }
            };

            var eventType = args.GetType().ToString().Split('.').Last();
            var eveEvent = new EveLogEvent
            {
                Event = eventType,
                Timestamp = args.Timestamp.ToTimestamp(),
                CharacterId = args.CharacterId,
                Amount = args.Amount,
                Pilot = args.Pilot,
                Ship = args.Ship,
                Weapon = args.Weapon,
                Application = args.Application,
                Corporation = args.Corporation,
                Alliance = args.Alliance,
            };
            try
            {
                await _client.PostEveLogEventAsync(eveEvent, meta);
                await _eventAggregator.PublishAsync(new LogStreamedEventArgs(args.CharacterId));
            }
            catch (RpcException ex)
            {
                if (ex.StatusCode == StatusCode.FailedPrecondition)
                {
                    Log.Debug("Server is not accepting logs, backing off for 30 seconds");
                    _lastBackOff = DateTimeOffset.Now;
                }
                else
                {
                    Log.Warning(ex, "Failed to send log event to server. Event: {@LogEvent}", eveEvent);
                }
            }
        });
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
    }
}