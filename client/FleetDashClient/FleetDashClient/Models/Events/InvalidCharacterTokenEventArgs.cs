namespace FleetDashClient.Models.Events;

public class InvalidCharacterTokenEventArgs : EventArgs
{
    public InvalidCharacterTokenEventArgs(string characterId)
    {
        CharacterId = characterId;
    }
    
    public string CharacterId { get; }
}