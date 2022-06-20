using Newtonsoft.Json;

namespace FleetDashClient.Models;

public class Token
{
    public string CharacterId { get; set; }
    
    [JsonProperty("access_token")]
    public string AccessToken { get; set; }
    
    [JsonProperty("refresh_token")]
    public string RefreshToken { get; set; }
    
    [JsonProperty("token_type")]
    public string TokenType { get; set; }
    
    [JsonProperty("expires_in")]
    public int ExpiresIn { get; set; }
    
    public DateTimeOffset ExpiresAt { get; set; }
    
    public virtual Character Character { get; set; }
}