namespace FleetDashClient.Models;

public class Character
{
    public string Id { get; set; }
    public string Name { get; set; }
    public string? OverviewPath { get; set; }
    
    public virtual ICollection<Token> Tokens { get; set; }
}