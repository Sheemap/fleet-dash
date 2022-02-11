using System.Collections.Concurrent;
using EventAggregator.Blazor;
using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;

namespace FleetDashClient.ViewModels;

public class IndexViewModel : IDisposable, IHandle<LogStreamedEventArgs>, IHandle<InvalidCharacterTokenEventArgs>, IHandle<CharacterTokenRenewedEventArgs>
{
    public List<Character> Characters { get; set; } = new();
    
    private readonly ConcurrentDictionary<string, (CharacterStatus,DateTimeOffset)> _lastLogStream = new();

    private readonly ICharacterService _characterService;
    private readonly IEventAggregator _eventAggregator;

    public IndexViewModel(ICharacterService characterService, IEventAggregator eventAggregator)
    {
        _characterService = characterService;
        _eventAggregator = eventAggregator;
        
        _eventAggregator.Subscribe(this);
        
        _characterService.OnCharacterAdded += HandleCharacterAdd;
        
        Characters = _characterService.GetCharacterListAsync().Result.ToList();
        
        UpdateStreamingStatus();
    }

    private Task UpdateStreamingStatus()
    {
        return Task.Run(async () =>
        {
            while (!CancellationToken.None.IsCancellationRequested)
            {
                await Task.Delay(5000);
                foreach (var (key, (prevStatus, prevTime)) in _lastLogStream)
                {
                    if (prevTime > DateTimeOffset.Now.AddSeconds(-30) && prevStatus != CharacterStatus.ActivelyStreaming)
                    {
                        await _eventAggregator.PublishAsync(
                            new CharacterStatusChangedEventArgs(key, CharacterStatus.ActivelyStreaming));
                        _lastLogStream[key] = (CharacterStatus.ActivelyStreaming, DateTimeOffset.Now);
                    }
                    else if(prevTime < DateTimeOffset.Now.AddSeconds(-30) && prevStatus == CharacterStatus.ActivelyStreaming)
                    {
                        await _eventAggregator.PublishAsync(
                            new CharacterStatusChangedEventArgs(key, CharacterStatus.Ready));
                        _lastLogStream[key] = (CharacterStatus.Ready, prevTime);

                    }
                }
            }
        });
    }
    
    public async Task HandleAsync(InvalidCharacterTokenEventArgs args)
    {
        await _eventAggregator.PublishAsync(
            new CharacterStatusChangedEventArgs(args.CharacterId, CharacterStatus.Error));
        _lastLogStream[args.CharacterId] = (CharacterStatus.Error, DateTimeOffset.MinValue);
    }
    
    public Task HandleAsync(LogStreamedEventArgs args)
    {
        _lastLogStream.AddOrUpdate(args.CharacterId, (CharacterStatus.ActivelyStreaming, DateTimeOffset.Now),
            (_, _) => (CharacterStatus.ActivelyStreaming, DateTimeOffset.Now));

        return _eventAggregator.PublishAsync(new CharacterStatusChangedEventArgs(args.CharacterId, CharacterStatus.ActivelyStreaming));
    }
        
    public async Task HandleAsync(CharacterTokenRenewedEventArgs args)
    {
        await _eventAggregator.PublishAsync(
            new CharacterStatusChangedEventArgs(args.CharacterId, CharacterStatus.Ready));
        _lastLogStream[args.CharacterId] = (CharacterStatus.Error, DateTimeOffset.MinValue);
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