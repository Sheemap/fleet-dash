using System.Diagnostics;
using System.IdentityModel.Tokens.Jwt;
using System.Net;
using System.Runtime.InteropServices;
using System.Text;
using FleetDashClient.Models;
using FleetDashClient.Services;
using FleetDashClient.Shared;
using Newtonsoft.Json;
using RestSharp;

namespace FleetDashClient.ViewModels;

public class EveLoginViewModel
{
    private const string ClientId = "NOPE";
    private const string ClientSecret = "NOPE";
    private const string RedirectUrl = "http://localhost:7845/eve-sso/callback";
    private const string TokenUrl = "https://login.eveonline.com/v2/oauth/token";
    private const string AuthUrlBase = "https://login.eveonline.com/v2/oauth/authorize";
    private readonly string[] Scopes = { "esi-fleets.read_fleet.v1", "esi-fleets.write_fleet.v1" };
    
    private readonly TimeSpan RequestTimeout = TimeSpan.FromSeconds(5);
    private readonly List<string> _validState = new();
    private CancellationTokenSource _cancellationTokenSource;
    private HttpListener _httpListener = new HttpListener();
    
    private readonly ICharacterService _characterService;

    public EveLoginViewModel(ICharacterService characterService)
    {
        _characterService = characterService;
        _httpListener.Prefixes.Add(RedirectUrl + "/");
    }
    
    public async Task StartAuthAsync()
        {
            var authUrl = new StringBuilder(AuthUrlBase);
            authUrl.Append("?response_type=code");
            authUrl.Append("&redirect_uri=" + RedirectUrl);
            authUrl.Append("&client_id=" + ClientId);
            authUrl.Append("&scope=" + string.Join("+", Scopes));
            authUrl.Append("&state=" + GenerateState());
            
    
            var url = authUrl.ToString();
            try
            {
                Process.Start(url);
            }
            catch
            {
                // hack because of this: https://github.com/dotnet/corefx/issues/10361
                if (RuntimeInformation.IsOSPlatform(OSPlatform.Windows))
                {
                    url = url.Replace("&", "^&");
                    Process.Start(new ProcessStartInfo("cmd", $"/c start {url}") { CreateNoWindow = true });
                }
                else if (RuntimeInformation.IsOSPlatform(OSPlatform.Linux))
                {
                    Process.Start("xdg-open", url);
                }
                else if (RuntimeInformation.IsOSPlatform(OSPlatform.OSX))
                {
                    Process.Start("open", url);
                }
                else
                {
                    throw;
                }
            }
    
            if (_cancellationTokenSource is { IsCancellationRequested: false })
            {
                _cancellationTokenSource.Cancel();
            }
            
            _cancellationTokenSource = new CancellationTokenSource(RequestTimeout);
    
            if (_httpListener.IsListening == false)
            {
                _httpListener.Start();
            }
            
            _httpListener.BeginGetContext(HandleCodeResponse, _httpListener);
    
            try
            {
                await Task.Delay(RequestTimeout, _cancellationTokenSource.Token);
            }
            catch(TaskCanceledException)
            {
            }
        }
    
        private string GenerateState()
        {
            var state = Guid.NewGuid().ToString();
            _validState.Add(state);
            
            // Remove the state after a minute
            Task.Run(async () =>
            {
                await Task.Delay(RequestTimeout);
                _validState.Remove(state);
            });
            
            return state;
        }
    
        private void HandleCodeResponse(IAsyncResult result)
        {
            result.AsyncWaitHandle.WaitOne();
            try
            {
                var request = _httpListener.EndGetContext(result);
                Task.Run(async () =>
                {
                    var code = request.Request.QueryString.GetValues("code")?[0];
                    var state = request.Request.QueryString.GetValues("state")?[0];
    
                    if (string.IsNullOrWhiteSpace(code) || string.IsNullOrWhiteSpace(state))
                    {
                        await SendResponseAsync("Unknown error processing request. Please try again.", request,
                            _cancellationTokenSource.Token);
                        return;
                    }
    
                    if (!_validState.Contains(state))
                    {
                        await SendResponseAsync("Invalid state param. Please try again.", request,
                            _cancellationTokenSource.Token);
                        return;
                    }
    
                    try
                    {
                        await GetTokenAsync(code);
                        await SendResponseAsync("Success! You may now close this window", request,
                            _cancellationTokenSource.Token);
                    }
                    catch (Exception ex)
                    {
                        await SendResponseAsync($"Unexpected error encountered! Please report this.</h1>" +
                                                $"<p>{ex.Message}</p><p>{ex.StackTrace}</p>", request,
                            _cancellationTokenSource.Token);
    
                    }
                }).Wait(_cancellationTokenSource.Token);
            }
            catch (Exception ex) when (ex is HttpListenerException or ObjectDisposedException)
            {
                // Ignore
            }
            
            // Cancel to indicate we are done
            // This is not the best way to do this, but Im tired of trying other ways :(
            // Might come revisit this when Im less tired
            _cancellationTokenSource.Cancel();
            
            // Also probs not the best way or place to do this but I dont wanna rn
            _httpListener.Stop();
        }
    
        private async Task SendResponseAsync(string message, HttpListenerContext context, CancellationToken token)
        {
            var htmlResponse = "<!DOCTYPE html><html><body><h1>" + message + "</h1></body></html>";
            var responseBytes = Encoding.UTF8.GetBytes(htmlResponse);
            await context.Response.OutputStream.WriteAsync(responseBytes, token);
            context.Response.Close();
        }
    
        private async Task GetTokenAsync(string code)
        {
            var client = new RestClient(TokenUrl);
            var request = new RestRequest();
            request.AddHeader("Content-Type", "application/x-www-form-urlencoded");
            request.AddParameter("grant_type", "authorization_code");
            request.AddParameter("code", code);
            request.AddParameter("redirect_uri", RedirectUrl);
            request.AddParameter("client_id", ClientId);
            request.AddParameter("client_secret", ClientSecret);
            var response = await client.ExecutePostAsync(request);
            if(response.StatusCode != HttpStatusCode.OK)
            {
                throw new Exception("Error getting token: " + response.Content);
            }
    
            var tokenSet = JsonConvert.DeserializeObject<Token>(response.Content);
            
            var token = new JwtSecurityToken(tokenSet.AccessToken);
            
            var characterSub = token.Payload["sub"].ToString();
            var splitSub = characterSub.Split("CHARACTER:EVE:");
            if (splitSub.Length != 2)
            {
                throw new Exception("Invalid character ID");
            }
    
            var characterName = token.Payload["name"].ToString();
            if (string.IsNullOrWhiteSpace(characterName))
            {
                throw new Exception("Invalid character name");
            }
            
            tokenSet.ExpiresAt = DateTimeOffset.Now + TimeSpan.FromSeconds(tokenSet.ExpiresIn);
    
            var dbChar = _characterService.GetCharacter(splitSub[1]);
            if (dbChar != null)
            {
                tokenSet.CharacterId = splitSub[1];
                _characterService.AddTokens(tokenSet);
            }
            else
            {
                var newCharacter = new Character
                {
                    Id = splitSub[1],
                    Name = characterName,
                    Tokens = new List<Token> { tokenSet },
                };
                await _characterService.AddCharacterAsync(newCharacter);
            }
        }
}