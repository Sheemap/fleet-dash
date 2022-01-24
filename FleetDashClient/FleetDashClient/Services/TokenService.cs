using FleetDashClient.Models;
using Newtonsoft.Json;
using RestSharp;

namespace FleetDashClient.Services;

public static class TokenService
{
    private const string ClientId = "NOPE";
    private const string ClientSecret = "NOPE";
    private const string TokenUrl = "https://login.eveonline.com/v2/oauth/token";

    public static async Task<Token> RefreshToken(Token token)
    {
        var client = new RestClient();
        var request = new RestRequest(TokenUrl, Method.Post);
        request.AddHeader("Content-Type", "application/x-www-form-urlencoded");
        request.AddParameter("grant_type", "refresh_token");
        request.AddParameter("refresh_token", token.RefreshToken);
        request.AddParameter("client_id", ClientId);
        request.AddParameter("client_secret", ClientSecret);
        var response = await client.ExecuteAsync(request);
        if (response.StatusCode != System.Net.HttpStatusCode.OK)
        {
            throw new Exception("Error refreshing token");
        }
        
        var content = JsonConvert.DeserializeObject<Token>(response.Content);
        if (content == null)
        {
            throw new Exception("Error refreshing token");
        }
        
        return content;
    }
}