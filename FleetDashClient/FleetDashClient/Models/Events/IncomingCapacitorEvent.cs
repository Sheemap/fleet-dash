namespace FleetDashClient.Models.Events;

public class IncomingCapacitorEvent : EveLogEventArgs
{
    public IncomingCapacitorEvent(string characterId, int amount, string pilot, string ship, string weapon,
        string application, string corporation, string alliance) :
        base(characterId, amount, pilot, ship, weapon, application, corporation, alliance)
    {
    }
}