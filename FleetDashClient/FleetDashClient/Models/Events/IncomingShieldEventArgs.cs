namespace FleetDashClient.Models.Events;

public class IncomingShieldEventArgs : EveLogEvent
{
    public IncomingShieldEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}