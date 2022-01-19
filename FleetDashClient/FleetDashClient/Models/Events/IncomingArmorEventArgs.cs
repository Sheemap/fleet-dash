namespace FleetDashClient.Models.Events;

public class IncomingArmorEventArgs : EveLogEvent
{
    public IncomingArmorEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}