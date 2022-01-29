namespace FleetDashClient.Models.Events;

public class LogStreamedEventArgs : EventArgs
{
    public LogStreamedEventArgs(string characterId)
    {
        CharacterId = characterId;
    }
    
    public string CharacterId { get; }
}