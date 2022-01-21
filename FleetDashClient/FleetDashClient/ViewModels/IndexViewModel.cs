using System.Diagnostics;
using System.Net;
using System.Runtime.InteropServices;
using System.Text;
using RestSharp;
using System;
using System.Collections.ObjectModel;
using System.IdentityModel.Tokens.Jwt;
using System.Reflection.Metadata.Ecma335;
using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using IdentityModel.OidcClient;
using Newtonsoft.Json;
using RestSharp.Authenticators;

namespace FleetDashClient.ViewModels;

public class IndexViewModel : IDisposable
{
    public List<Character> Characters { get; set; } = new();

    private readonly ICharacterService _characterService;

    public IndexViewModel(ICharacterService characterService)
    {
        _characterService = characterService;
        _characterService.OnCharacterAdded += HandleCharacterAdd;
        
        Characters = _characterService.GetCharacterListAsync().Result.ToList();
    }
    
    public Task RemoveCharacterAsync(string characterId)
    {
        Characters.RemoveAll(x => x.Id == characterId);
        return _characterService.RemoveCharacterAsync(characterId);
    }

    private void HandleCharacterAdd(object sender, CharacterAddedEventArgs args)
    {
        Characters.Add(args.Character);
    }

    public void Dispose()
    {
        _characterService.OnCharacterAdded -= HandleCharacterAdd;
    }
}