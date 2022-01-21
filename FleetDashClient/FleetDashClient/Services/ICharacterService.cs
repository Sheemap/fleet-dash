using FleetDashClient.Models;
using FleetDashClient.Models.Events;

namespace FleetDashClient.Services;

public interface ICharacterService
{
    Task<IEnumerable<Character>> GetCharacterListAsync();
    Task<Character> AddCharacterAsync(Character character);
    CharacterStatus GetCharacterStatus(string characterId);
    Task RemoveCharacterAsync(string characterId);
    Character? GetCharacter(string characterId);
    void AddTokens(params Token[] tokens);
    event EventHandler<CharacterAddedEventArgs> OnCharacterAdded;
    event EventHandler<CharacterRemovedEventArgs> OnCharacterRemoved;
}