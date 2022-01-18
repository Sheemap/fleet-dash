namespace FleetDashClient.Constants;

public static class EnglishRegex
{
    public const string PilotAndWeapon = @"(?:.*ffffffff>(?<Name>[^\(\)<>]*)\[(?<Tag>.*)\](?:.*\((?<Ship>.*)\)<|<)/b.*> \-(?: (?<Weapon>.*?)(?: \-|<)|.*)) (?<Application>[A-z]*$)";
    public const string IncomingDamage = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*>from<";
    public const string OutgoingDamage = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*>to<";

}