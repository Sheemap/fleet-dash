using FleetDashClient.Models.Events;

namespace FleetDashClient.Services;

public interface ILogParserService
{
    public event EventHandler<IncomingDamageEvent> OnIncomingDamage;
    public event EventHandler<OutgoingDamageEvent> OnOutgoingDamage;
    event EventHandler<OutgoingArmorEvent> OnOutgoingArmor;
    event EventHandler<IncomingArmorEvent> OnIncomingArmor;
    event EventHandler<IncomingCapacitorEvent> OnIncomingCapacitor;
    event EventHandler<OutgoingCapacitorEvent> OnOutgoingCapacitor;
    event EventHandler<IncomingHullEvent> OnIncomingHull;
    event EventHandler<OutgoingHullEvent> OnOutgoingHull;
    event EventHandler<IncomingShieldEvent> OnIncomingShield;
    event EventHandler<OutgoingShieldEvent> OnOutgoingShield;
    event EventHandler<IncomingNeutEvent> OnIncomingNeut;
    event EventHandler<OutgoingNeutEvent> OnOutgoingNeut;
    event EventHandler<IncomingNosEvent> OnIncomingNos;
    event EventHandler<OutgoingNosEvent> OnOutgoingNos;
    List<string> WatchedCharacters { get; }
    void StartWatchingCharacter(string characterId, string overviewSettings);
    void StopWatchingCharacter(string characterId);
}