namespace FleetDashClient.Models.Events;

public class OutgoingHullEventArgs : EveLogEvent
{
    public OutgoingHullEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}