using FleetDashClient.Models;

namespace FleetDashClient.Services;

public interface ICharacterService
{
    Task<IEnumerable<Character>> GetCharacterListAsync();
    Task<Character> AddCharacterAsync(Character character);
    CharacterStatus GetCharacterStatus(string characterId);
    Task RemoveCharacterAsync(string characterId);
    Character? GetCharacter(string characterId);
    void AddTokens(params Token[] tokens);
}