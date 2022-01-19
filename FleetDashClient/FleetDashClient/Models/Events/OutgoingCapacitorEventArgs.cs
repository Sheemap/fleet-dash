namespace FleetDashClient.Models.Events;

public class OutgoingCapacitorEventArgs : EveLogEvent
{
    public OutgoingCapacitorEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}