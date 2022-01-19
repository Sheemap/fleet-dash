namespace FleetDashClient.Models.Events;

public class IncomingHullEventArgs : EveLogEvent
{
    public IncomingHullEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}