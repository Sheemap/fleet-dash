namespace EveLogParser.Constants;

public static class EnglishRegex
{
    public const string Timestamp = @"\[ (?<timestamp>\d{4}\.\d{2}\.\d{2} \d{2}:\d{2}:\d{2}) \] ";

    // These are static throughout overview settings, so we can just mash these in
    public const string IncomingDamage = @"\[ (?<timestamp>\d{4}\.\d{2}\.\d{2} \d{2}:\d{2}:\d{2}) \] \(combat\) <color=0xffcc0000><b>(?<amount>[0-9]+)</b> <color=0x77ffffff><font size=10>from</font> <b><color=0xffffffff>(?<pilotname>[A-Za-z0-9\- ']+)\[(?<corporation>[A-Za-z0-9\- ']+)]\((?<shiptype>[A-Za-z0-9\- ']+)\)</b><font size=10><color=0x77ffffff> - (?<weapon>[A-Za-z0-9\- ']+) - (?<application>[A-Za-z0-9\- ']+)";
    public const string OutgoingDamage = @"\[ (?<timestamp>\d{4}\.\d{2}\.\d{2} \d{2}:\d{2}:\d{2}) \] \(combat\) <color=0xff00ffff><b>(?<amount>[0-9]+)</b> <color=0x77ffffff><font size=10>to</font> <b><color=0xffffffff>(?<pilotname>[A-Za-z0-9\- ']+)\[(?<corporation>[A-Za-z0-9\- ']+)]\((?<shiptype>[A-Za-z0-9\- ']+)\)</b><font size=10><color=0x77ffffff> - (?<weapon>[A-Za-z0-9\- ']+) - (?<application>[A-Za-z0-9\- ']+)";

    // Append this to any of the other events to grab the weapon used
    public const string Weapon = @".*?<color=0x77ffffff><font size=10> - (?<weapon>[A-Za-z0-9\- \']+)</font>";

    // This regex matches any log entries we dont want to parse, which helps us decrease the amount of logs to parse, and increase performance
    public const string NoParse = @"((?:\(mining\))|(?:\(notify\)))|(?:Warp ((?:disruption)|(?:scramble)) attempt)";

    public const string IncomingHull = @"\(combat\).*?<b>(?<amount>[0-9]+).*?remote hull repaired by";
    public const string OutgoingHull = @"\(combat\).*?<b>(?<amount>[0-9]+).*?remote hull repaired to";
    public const string IncomingArmor = @"\(combat\).*?<b>(?<amount>[0-9]+).*?remote armor repaired by";
    public const string OutgoingArmor = @"\(combat\).*?<b>(?<amount>[0-9]+).*?remote armor repaired to";
    public const string IncomingShield = @"\(combat\).*?<b>(?<amount>[0-9]+).*?remote shield boosted by";
    public const string OutgoingShield = @"\(combat\).*?<b>(?<amount>[0-9]+).*?remote shield boosted to";
    public const string IncomingCapacitor = @"\(combat\) .*?<b>(?<amount>[0-9]+).*?remote capacitor transmitted by";
    public const string OutgoingCapacitor = @"\(combat\) .*?<b>(?<amount>[0-9]+).*?remote capacitor transmitted to";
    public const string IncomingNeut = @"\(combat\).*?ffe57f7f><b>(?<amount>[0-9]+).*?energy neutralized";
    public const string OutgoingNeut = @"\(combat\).*?ff7fffff><b>(?<amount>[0-9]+).*?energy neutralized";
    public const string IncomingNos = @"\(combat\).*?<b>\-(?<amount>[0-9]+).*?energy drained to";
    public const string OutgoingNos = @"\(combat\).*?<b>\+(?<amount>[0-9]+).*?energy drained from";
    public const string IncomingJam = @"\(combat\) <color=0x77ffffff><font size=10>You're</font> <color=0xffffffff><b>jammed</b> <color=0x77ffffff><font size=10>by</font>";
    public const string OutgoingJam = @"\(combat\) <color=0xffffffff><b>.*? jammed";
}