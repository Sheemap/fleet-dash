namespace FleetDashClient.Models.Events;

public class IncomingDamageEvent : EveLogEventArgs
{
    public IncomingDamageEvent(string characterId, int amount, string pilot, string ship, string weapon,
        string application, string corporation, string alliance) : base(characterId, amount, pilot, ship, weapon,
        application, corporation, alliance)
    {
    }
}