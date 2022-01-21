namespace FleetDashClient.Models.Events;

public class OutgoingShieldEvent : EveLogEventArgs
{
    public OutgoingShieldEvent(string characterId, int amount, string pilot, string ship, string weapon,
        string application, string corporation, string alliance) :
        base(characterId, amount, pilot, ship, weapon, application, corporation, alliance)
    {
    }
}