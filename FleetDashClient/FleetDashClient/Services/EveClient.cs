using System.Net;
using Newtonsoft.Json;
using RestSharp;
using Serilog;

namespace FleetDashClient.Services;

public static class EveClient
{
    public static async Task<string?> GetCharacterFleetIdAsync(SemaphoreSlim semaphore, string characterId, string token)
    {
        Log.Debug("Fetching character fleet");
        try
        {
            // TODO: Return cached value if it exists, GrpcLogShipper should not be in charge of that
            if (!await semaphore.WaitAsync(5000))
            {
                throw new TimeoutException("Timed out waiting for semaphore");
            }
            var client = new RestClient("https://esi.evetech.net/latest");
            var request = new RestRequest("/characters/{character_id}/fleet/", Method.Get);
            request.AddHeader("Authorization", $"Bearer {token}");
            request.AddUrlSegment("character_id", characterId);

            var response = await client.ExecuteGetAsync(request);
            if ((int) response.StatusCode >= 400)
            {
                Log.Warning("Error occurred fetching fleet. {StatusCode} - {ResponseContent}", response.StatusCode,
                    response.Content);
                return null;
            }

            var fleet = JsonConvert.DeserializeObject<dynamic>(response.Content);
            Log.Debug("Fleet fetched {@CharacterFleet}", fleet.fleet_id);
            return fleet.fleet_id.ToString();
        }
        finally
        {
            semaphore.Release();
        }
    }
    
    public static async Task<int?> GetCharacterShipTypeId(SemaphoreSlim semaphore, string characterId, string token)
    {
        Log.Debug("Fetching character ship info");
        try
        {
            // TODO: Return cached value if it exists, GrpcLogShipper should not be in charge of that
            if (!await semaphore.WaitAsync(5000))
            {
                throw new TimeoutException("Timed out waiting for semaphore");
            }
            var client = new RestClient("https://esi.evetech.net/latest");
            var request = new RestRequest("/characters/{character_id}/ship/", Method.Get);
            request.AddHeader("Authorization", $"Bearer {token}");
            request.AddUrlSegment("character_id", characterId);

            var response = await client.ExecuteGetAsync(request);
            if ((int) response.StatusCode >= 400)
            {
                Log.Warning("Error occurred fetching ship info. {StatusCode} - {ResponseContent}", response.StatusCode,
                    response.Content);
                return null;
            }

            var ship = JsonConvert.DeserializeObject<dynamic>(response.Content);
            Log.Debug("Ship info fetched. ShipTypeId is {@CharacterShipTypeId}", ship.ship_type_id);
            return ship.ship_type_id;
        }
        finally
        {
            semaphore.Release();
        }
    }
}