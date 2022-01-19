namespace FleetDashClient.Constants;

public static class EnglishRegex
{
    public const string PilotAndWeapon = @"(?:.*ffffffff>(?<Name>[^\(\)<>]*)\[(?<Tag>.*)\](?:.*\((?<Ship>.*)\)<|<)/b.*> \-(?: (?<Weapon>.*?)(?: \-|<)|.*)) (?<Application>[A-z]*$)";
    public const string IncomingDamage = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*>from<";
    public const string OutgoingDamage = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*>to<";
    public const string IncomingHull = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote hull repaired by <";
    public const string OutgoingHull = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote hull repaired to <";
    public const string IncomingArmor = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote armor repaired by <";
    public const string OutgoingArmor = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote armor repaired to <";
    public const string IncomingShield = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote shield boosted by <";
    public const string OutgoingShield = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote shield boosted to <";
    public const string IncomingCapacitor = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote capacitor transmitted by <";
    public const string OutgoingCapacitor = @"\(combat\) <.*?><b>(?<Amount>[0-9]+).*> remote capacitor transmitted to <";
    public const string IncomingNeut = @"\(combat\) <.*?ffe57f7f><b>(?<Amount>[0-9]+).*> energy neutralized <";
    public const string OutgoingNeut = @"\(combat\) <.*?ff7fffff><b>(?<Amount>[0-9]+).*> energy neutralized <";
    public const string IncomingNos = @"\(combat\) <.*?><b>\-(?<Amount>[0-9]+).*> energy drained to <";
    public const string OutgoingNos = @"\(combat\) <.*?><b>\+(?<Amount>[0-9]+).*> energy drained from <";
}