using ElectronNET.API;

namespace FleetDashClient.Models;

public class Configuration
{
    public int Id { get; set; }
    public string LogDirectory { get; set; }
    public string? Overview { get; set; }
    public int WindowHeight { get; set; }
    public int WindowWidth { get; set; }
    public int WindowPositionX { get; set; }
    public int WindowPositionY { get; set; }
}