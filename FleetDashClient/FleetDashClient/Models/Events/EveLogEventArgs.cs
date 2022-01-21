namespace FleetDashClient.Models.Events;

public class EveLogEventArgs : EventArgs
{
    public EveLogEventArgs(string characterId, int amount, string pilot, string ship, string weapon, string application,
        string corporation, string alliance)
    {
        CharacterId = characterId;
        Amount = amount;
        Pilot = pilot;
        Ship = ship;
        Weapon = weapon;
        Application = application;
        Corporation = corporation;
        Alliance = alliance;
    }

    public string CharacterId { get; }
    public int Amount { get; }
    public string Pilot { get; }
    public string Ship { get; }
    public string Weapon { get; }
    public string Application { get; }
    public string Corporation { get; }
    public string Alliance { get; }
}