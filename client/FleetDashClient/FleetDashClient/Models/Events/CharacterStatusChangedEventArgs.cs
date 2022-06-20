namespace FleetDashClient.Models.Events;

public class CharacterStatusChangedEventArgs
{
    public CharacterStatusChangedEventArgs(string characterId, CharacterStatus status)
    {
        CharacterId = characterId;
        Status = status;
    }

    public string CharacterId { get; }
    public CharacterStatus Status { get; }
}