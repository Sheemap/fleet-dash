using EveLogParser.Models.Events;

namespace EveLogParser;

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
    event EventHandler<IncomingJamEvent>? OnIncomingJam;
    event EventHandler<OutgoingJamEvent>? OnOutgoingJam;
    void StartWatchingCharacter(string characterId, string? overviewPath);
    void StopWatchingCharacter(string characterId);
    Task Start();
    void Stop();
}