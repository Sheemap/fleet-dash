namespace FleetDashClient.Models.Events;

public class PathUpdatedEventArgs : EventArgs
{
    public PathUpdatedEventArgs(string path)
    {
        Path = path;
    }
    
    public string Path { get; }
}