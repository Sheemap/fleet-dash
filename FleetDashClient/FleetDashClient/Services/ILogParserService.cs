using FleetDashClient.Models.Events;

namespace FleetDashClient.Services;

public interface ILogParserService
{
    public event EventHandler<IncomingDamageEventArgs> RaiseIncomingDamageEvent;
    public event EventHandler<OutgoingDamageEventArgs> RaiseOutgoingDamageEvent;
    event EventHandler<OutgoingArmorEventArgs> RaiseOutgoingArmorEvent;
    event EventHandler<IncomingArmorEventArgs> RaiseIncomingArmorEvent;
    event EventHandler<IncomingCapacitorEventArgs> RaiseIncomingCapacitorEvent;
    event EventHandler<OutgoingCapacitorEventArgs> RaiseOutgoingCapacitorEvent;
    event EventHandler<IncomingHullEventArgs> RaiseIncomingHullEvent;
    event EventHandler<OutgoingHullEventArgs> RaiseOutgoingHullEvent;
    event EventHandler<IncomingShieldEventArgs> RaiseIncomingShieldEvent;
    event EventHandler<OutgoingShieldEventArgs> RaiseOutgoingShieldEvent;
    event EventHandler<IncomingNeutEventArgs> RaiseIncomingNeutEvent;
    event EventHandler<OutgoingNeutEventArgs> RaiseOutgoingNeutEvent;
    event EventHandler<IncomingNosEventArgs> RaiseIncomingNosEvent;
    event EventHandler<OutgoingNosEventArgs> RaiseOutgoingNosEvent;
    void StartWatchingCharacter(string characterId, string overviewSettings);
    void StopWatchingCharacter(string characterId);
}