namespace FleetDashClient.Models.Events;

public class IncomingDamageEventArgs : EveLogEvent
{
    public IncomingDamageEventArgs(string characterId, int amount, string pilot, string ship, string weapon,
        string application) : base(characterId, amount, pilot, ship, weapon)
    {
        Application = application;
    }

    public string Application { get; }
}