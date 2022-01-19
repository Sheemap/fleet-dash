namespace FleetDashClient.Models.Events;

public class IncomingNosEventArgs : EveLogEvent
{
    public IncomingNosEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}