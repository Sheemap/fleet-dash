using EveLogParser;
using EveLogParser.Models.Events;
using FleetDashClient.Data;
using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using FleetDashClient.Protobuf;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Grpc.Net.Client;
using Microsoft.EntityFrameworkCore;
using Quobject.SocketIoClientDotNet.Client;

namespace FleetDashClient.Services;

public class GrpcLogShipper
{
    private readonly ILogParserService _logParserService;
    private readonly DataContext _dbContext;
    
    private (string?, DateTimeOffset) _fleetData;
    private DateTimeOffset _lastBackOff = DateTimeOffset.MinValue;
    
    private readonly FleetDashService.FleetDashServiceClient _client;
    
    public GrpcLogShipper(ILogParserService logParserService, DataContext dataContext)
    {
        _logParserService = logParserService;
        _dbContext = dataContext;

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

        var channel = GrpcChannel.ForAddress("http://localhost:50051");
        _client = new FleetDashService.FleetDashServiceClient(channel);
    }

    private async Task<Token> GetFreshToken(Token token, string characterId)
    {
        if (token.ExpiresAt > DateTimeOffset.Now.AddMinutes(1)) return token;
        
        // refresh token
        token = await TokenService.RefreshToken(token);
        token.CharacterId = characterId;
        token.ExpiresAt = DateTimeOffset.Now.AddSeconds(token.ExpiresIn);
        _dbContext.Tokens.Add(token);
        await _dbContext.SaveChangesAsync();

        return token;
    }

    private void EveLogHandler<T>(object? sender, T args) where T : EveLogEventArgs
    {
        if (_lastBackOff < DateTimeOffset.Now.AddSeconds(-30))
        {
            return;
        }

        Task.Run(async () =>
        {
            var character = await _dbContext.Characters
                .Include(x => x.Tokens)
                .FirstOrDefaultAsync(x => x.Id == args.CharacterId);
            if (character == null)
            {
                return;
            }
            
            var token = character.Tokens.OrderByDescending(x => x.ExpiresAt).FirstOrDefault();
            if (token == null)
            {
                // TODO: Inform UI that character has no token
                return;
            }
            
            token = await GetFreshToken(token, character.Id);

            // Need to be in a fleet to ship logs
            if (_fleetData.Item2 < DateTimeOffset.Now.AddSeconds(-30))
            {
                var fleetId = await EveClient.GetCharacterFleetIdAsync(token.CharacterId, token.AccessToken);
                _fleetData = (fleetId, DateTimeOffset.Now);
            }
            if (_fleetData.Item1 == null)
            {
                return;
            }

            var meta = new Metadata
            {
                { "Authorization", $"Bearer {token.AccessToken}" }
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
            }
            catch (RpcException ex)
            {
                if (ex.StatusCode == StatusCode.FailedPrecondition)
                {
                    _lastBackOff = DateTimeOffset.Now;
                }
            }
        });
    }
}