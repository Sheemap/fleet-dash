﻿using System.Net;
using Newtonsoft.Json;
using RestSharp;

namespace FleetDashClient.Services;

public class EveClient
{
    public static async Task<string?> GetCharacterFleetIdAsync(string characterId, string token)
    {
        var client = new RestClient("https://esi.evetech.net/latest");
        var request = new RestRequest("/characters/{character_id}/fleet/", Method.Get);
        request.AddUrlSegment("character_id", characterId);
        
        var response = await client.ExecuteGetAsync(request);
        if (response.StatusCode != HttpStatusCode.OK) return null;
        
        var fleet = JsonConvert.DeserializeObject<dynamic>(response.Content);
        return fleet.fleet_id.ToString();

    }
}