using FleetDashClient.Data;
using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using Microsoft.EntityFrameworkCore;

namespace FleetDashClient.Services;

public class CharacterService : ICharacterService
{
    public event EventHandler<CharacterAddedEventArgs> OnCharacterAdded;
    public event EventHandler<CharacterModifiedEventArgs> OnCharacterModified;
    public event EventHandler<CharacterRemovedEventArgs> OnCharacterRemoved;
    
    private readonly DataContext _dbContext;
    
    public CharacterService(DataContext dbContext)
    {
        _dbContext = dbContext;
    }
    
    public Task<IEnumerable<Character>> GetCharacterListAsync()
    {
        return Task.FromResult(_dbContext.Characters.AsEnumerable());
    }
    
    public async Task<Character> AddCharacterAsync(Character character)
    {
        _dbContext.Characters.Add(character);
        await _dbContext.SaveChangesAsync();

        var raiseEvent = OnCharacterAdded;
        if (raiseEvent != null) raiseEvent(this, new CharacterAddedEventArgs(character));
        
        return character;
    }
    
    public async Task ModifyCharacterAsync(string characterId, string? overviewPath)
    {
        var character = _dbContext.Characters.FirstOrDefault(x => x.Id == characterId);
        if (character != null)
        {
            character.OverviewPath = overviewPath;
            await _dbContext.SaveChangesAsync();
        }

        var raiseEvent = OnCharacterModified;
        if (raiseEvent != null) raiseEvent(this, new CharacterModifiedEventArgs(characterId, overviewPath));
    }

    public async Task RemoveCharacterAsync(string characterId)
    {
        var character = _dbContext.Characters.FirstOrDefault(x => x.Id == characterId);
        if (character != null)
            _dbContext.Characters.Remove(character);

        var tokens = _dbContext.Tokens.Where(x => x.CharacterId == characterId);
        _dbContext.Tokens.RemoveRange(tokens);

        var raiseEvent = OnCharacterRemoved;
        if (raiseEvent != null) raiseEvent(this, new CharacterRemovedEventArgs(character));
        
        await _dbContext.SaveChangesAsync();
    }
    
    public Character? GetCharacter(string characterId)
    {
        return _dbContext.Characters.FirstOrDefault(x => x.Id == characterId);
    }

    public CharacterStatus GetCharacterStatus(string characterId)
    {
        var character = _dbContext.Characters
            .Include(x => x.Tokens)
            .FirstOrDefault(x => x.Id == characterId);
        if (character == null)
        {
            return CharacterStatus.Error;
        }
        
        return character.Tokens.Any(x => x.ExpiresAt > DateTimeOffset.Now)
            ? CharacterStatus.Ready
            : CharacterStatus.Error;
    }

    public void AddTokens(params Token[] tokens)
    {
        _dbContext.Tokens.AddRange(tokens);
        _dbContext.SaveChanges();
    }
}