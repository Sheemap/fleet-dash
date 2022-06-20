namespace FleetDashClient.Models.Events;

public class CharacterTokenRenewedEventArgs : EventArgs
{
    public CharacterTokenRenewedEventArgs(string characterId)
    {
        CharacterId = characterId;
    }
    
    public string CharacterId { get; }
}