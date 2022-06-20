namespace FleetDashClient.Models.Events;

public class CharacterAddedEventArgs : EventArgs
{
    public CharacterAddedEventArgs(Character character)
    {
        Character = character;
    }
    
    public Character Character { get; }
}