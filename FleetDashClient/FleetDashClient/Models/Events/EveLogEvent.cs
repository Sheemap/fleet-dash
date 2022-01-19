namespace FleetDashClient.Models.Events;

public class EveLogEvent : EventArgs
{
    public EveLogEvent(string characterId, int amount, string pilot, string ship, string weapon)
    {
        CharacterId = characterId;
        Amount = amount;
        Pilot = pilot;
        Ship = ship;
        Weapon = weapon;
    }

    public string CharacterId { get; }
    public int Amount { get; }
    public string Pilot { get; }
    public string Ship { get; }
    public string Weapon { get; }
}