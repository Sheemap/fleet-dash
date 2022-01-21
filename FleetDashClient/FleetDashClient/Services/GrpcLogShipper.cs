using FleetDashClient.Models.Events;
using FleetDashClient.Protobuf;
using Google.Protobuf.WellKnownTypes;
using Grpc.Net.Client;

namespace FleetDashClient.Services;

public class GrpcLogShipper
{
    private readonly ILogParserService _logParserService;
    private readonly FleetDashService.FleetDashServiceClient _client;
    
    public GrpcLogShipper(ILogParserService logParserService)
    {
        _logParserService = logParserService;

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

    private void EveLogHandler<T>(object sender, T args) where T : EveLogEventArgs
    {
        Task.Run(async () =>
        {
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
            await _client.PostEveLogEventAsync(eveEvent);
        });
    }
}