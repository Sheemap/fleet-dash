namespace FleetDashClient.Models.Events;

public class IncomingNeutEventArgs : EveLogEvent
{
    public IncomingNeutEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}