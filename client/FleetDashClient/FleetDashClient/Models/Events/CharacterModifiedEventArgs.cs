namespace FleetDashClient.Models.Events;

public class CharacterModifiedEventArgs : EventArgs
{
    public CharacterModifiedEventArgs(string characterId, string? overviewPath)
    {
        CharacterId = characterId;
        OverviewPath = overviewPath;
    }
    
    public string CharacterId { get; }
    public string? OverviewPath { get; }
}