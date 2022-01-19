namespace FleetDashClient.Models.Events;

public class LogFileReadEventArgs : EventArgs
{
    public LogFileReadEventArgs(string characterId, byte[] content)
    {
        CharacterId = characterId;
        Content = content;
    }

    public string CharacterId { get; set; }
    public byte[] Content { get; set; }
}