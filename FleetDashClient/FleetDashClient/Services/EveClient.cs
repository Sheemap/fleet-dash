using System.Net;
using Newtonsoft.Json;
using RestSharp;

namespace FleetDashClient.Services;

public static class EveClient
{
    public static async Task<string?> GetCharacterFleetIdAsync(string characterId, string token)
    {
        var client = new RestClient("https://esi.evetech.net/latest");
        var request = new RestRequest("/characters/{character_id}/fleet/", Method.Get);
        request.AddHeader("Authorization", $"Bearer {token}");
        request.AddUrlSegment("character_id", characterId);
        
        var response = await client.ExecuteGetAsync(request);
        if (response.StatusCode != HttpStatusCode.OK) return null;
        
        var fleet = JsonConvert.DeserializeObject<dynamic>(response.Content);
        return fleet.fleet_id.ToString();

    }
    
    public static async Task<int?> GetCharacterShipTypeId(string characterId, string token)
    {
        var client = new RestClient("https://esi.evetech.net/latest");
        var request = new RestRequest("/characters/{character_id}/ship/", Method.Get);
        request.AddHeader("Authorization", $"Bearer {token}");
        request.AddUrlSegment("character_id", characterId);
        
        var response = await client.ExecuteGetAsync(request);
        if (response.StatusCode != HttpStatusCode.OK) return null;
        
        var ship = JsonConvert.DeserializeObject<dynamic>(response.Content);
        return ship.ship_type_id;

    }
}