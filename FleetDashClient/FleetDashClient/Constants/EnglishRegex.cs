namespace FleetDashClient.Constants;

public static class EnglishRegex
{
    public const string PilotAndWeapon = @"(?:.*ffffffff>(?<Name>[^\(\)<>]*)\[(?<Tag>.*)\](?:.*\((?<Ship>.*)\)<|<)/b.*> \-(?: (?<Weapon>.*?)(?: \-|<)|.*)) (?<Application>[A-z]*$)";
    public const string IncomingDamage = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*>from<";
    public const string OutgoingDamage = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*>to<";
    public const string IncomingArmor = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote armor repaired by <";
    public const string OutgoingArmor = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote armor repaired to <";
}