namespace FleetDashClient.Models.Events;

public class CharacterRemovedEventArgs : EventArgs
{
    public CharacterRemovedEventArgs(Character character)
    {
        Character = character;
    }
    
    public Character Character { get; }
}