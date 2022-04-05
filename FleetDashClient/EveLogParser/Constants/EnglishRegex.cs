namespace EveLogParser.Constants;

public static class EnglishRegex
{
    public const string Timestamp = @"\[ (?<Timestamp>\d{4}\.\d{2}\.\d{2} \d{2}:\d{2}:\d{2}) \] ";

    // Used for damage, and for when no overview is set
    public const string DefaultShipLabels =
        @"</font> <b><color=0xffffffff>(?<ShipLabel>[A-Z|a-z|0-9|\-| \']+)(\[(?<ShipLabel>[A-Z|a-z|0-9|\-| \']+)\])?\((?<ShipLabel>[A-Z|a-z|0-9|\-| \']+)\)</b><font size=10><color=0x77ffffff> - (?<ShipLabel>[A-Z|a-z|0-9|\-| \']+)";
    
    // The idea of this regex is to match everything that is potentially a name for something
    // We put this after the above regexs, so we can only find names after the initial match
    // Then we sort the matches according to the shipLabelOrder, look for missing values, and do our best to place things in the right spot
    // This will be the most robust way that works across essentially any overview settings
    public const string ShipLabelGroups =
        @"(.*?[^(A-Z|a-z|0-9|\-|'| |<|\])]+(?<ShipLabel>[A-Z|a-z|0-9|\-|'| ]{3,37})[^(A-Z|a-z|0-9|\-|'| |>|[)]+.*?)*";

    // This regex matches any log entries we dont want to parse, which helps us decrease the amount of logs to parse, and increase performance
    public const string NoParse = @"((?:\(mining\))|(?:\(notify\)))|(?:Warp ((?:disruption)|(?:scramble)) attempt)";

    public const string IncomingDamage = @"\(combat\).*?<b>(?<Amount>[0-9]+).*?from";
    public const string OutgoingDamage = @"\(combat\).*?<b>(?<Amount>[0-9]+).*?to";
    public const string IncomingHull = @"\(combat\).*?<b>(?<Amount>[0-9]+).*?remote hull repaired by";
    public const string OutgoingHull = @"\(combat\).*?<b>(?<Amount>[0-9]+).*?remote hull repaired to";
    public const string IncomingArmor = @"\(combat\).*?<b>(?<Amount>[0-9]+).*?remote armor repaired by";
    public const string OutgoingArmor = @"\(combat\).*?<b>(?<Amount>[0-9]+).*?remote armor repaired to";
    public const string IncomingShield = @"\(combat\).*?<b>(?<Amount>[0-9]+).*?remote shield boosted by";
    public const string OutgoingShield = @"\(combat\).*?<b>(?<Amount>[0-9]+).*?remote shield boosted to";

    public const string IncomingCapacitor =
        @"\(combat\) .*?<b>(?<Amount>[0-9]+).*?remote capacitor transmitted by";

    public const string OutgoingCapacitor =
        @"\(combat\) .*?<b>(?<Amount>[0-9]+).*?remote capacitor transmitted to";

    public const string IncomingNeut = @"\(combat\).*?ffe57f7f><b>(?<Amount>[0-9]+).*?energy neutralized";
    public const string OutgoingNeut = @"\(combat\).*?ff7fffff><b>(?<Amount>[0-9]+).*?energy neutralized";
    public const string IncomingNos = @"\(combat\).*?<b>\-(?<Amount>[0-9]+).*?energy drained to";
    public const string OutgoingNos = @"\(combat\).*?<b>\+(?<Amount>[0-9]+).*?energy drained from";

    public const string IncomingJam = @"\(combat\).*?You're.*?jammed.*?by";

    public const string OutgoingJam = @"\(combat\).*?(.*?[^(A-Z|a-z|0-9|\-|'| |<|\])]+(?<ShipLabel>[A-Z|a-z|0-9|\-|'| ]{3,37})[^(A-Z|a-z|0-9|\-|'| |>|[)]+)+.* - jammed.* - (?<ShipLabel>[A-Z|a-z|0-9|\-|'| ]+)<";
}