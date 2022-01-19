namespace FleetDashClient.Models.Events;

public class OutgoingShieldEventArgs : EveLogEvent
{
    public OutgoingShieldEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}