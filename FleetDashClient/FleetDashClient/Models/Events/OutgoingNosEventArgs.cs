namespace FleetDashClient.Models.Events;

public class OutgoingNosEventArgs : EveLogEvent
{
    public OutgoingNosEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}