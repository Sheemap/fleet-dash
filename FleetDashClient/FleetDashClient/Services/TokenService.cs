using FleetDashClient.Models;
using FleetDashClient.Models.Exceptions;
using Newtonsoft.Json;
using RestSharp;

namespace FleetDashClient.Services;

public static class TokenService
{
    private const string ClientId = "f5359c448f2c444d960d1ce8d3047397";
    private const string ClientSecret = "SMBDOxOp0flzin3bO0YObnoCVlaZ22ofBogIevwl";
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
            throw new TokenRefreshException($"Received {response.StatusCode}. Content: {response.Content}");
        }
        
        var content = JsonConvert.DeserializeObject<Token>(response.Content);
        if (content == null)
        {
            throw new TokenRefreshException("Failed to parse the response");
        }
        
        content.ExpiresAt = DateTimeOffset.Now.AddSeconds(content.ExpiresIn);

        return content;
    }
}