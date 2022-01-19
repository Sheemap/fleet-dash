namespace FleetDashClient.Models.Events;

public class OutgoingNeutEventArgs : EveLogEvent
{
    public OutgoingNeutEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}