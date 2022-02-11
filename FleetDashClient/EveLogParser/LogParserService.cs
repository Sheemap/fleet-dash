using System.Globalization;
using System.Text;
using System.Text.RegularExpressions;
using EveLogParser.Builder;
using EveLogParser.Constants;
using EveLogParser.Models.Events;
using Microsoft.Extensions.Options;
using Serilog;
using YamlDotNet.Core;
using YamlDotNet.Serialization;

namespace EveLogParser;

public class LogParserService : ILogParserService, IDisposable
{
    private readonly List<string> _watchedCharacters = new();
    private string? _overviewRegex;
    
    private readonly ILogReaderService _logReaderService;
    
    public Task Start() => _logReaderService.Start();
    public void Stop() => _logReaderService.Stop();

    public LogParserService(ILogReaderService logReaderService, IOptionsMonitor<EveLogParserOptions> options)
    {
        _logReaderService = logReaderService;
        _logReaderService.OnFileRead += HandleLogFileRead;
        
        UpdateRegex(options.CurrentValue);
        options.OnChange(UpdateRegex);
    }

    public event EventHandler<IncomingDamageEvent>? OnIncomingDamage;
    public event EventHandler<OutgoingDamageEvent>? OnOutgoingDamage;
    public event EventHandler<IncomingHullEvent>? OnIncomingHull;
    public event EventHandler<OutgoingHullEvent>? OnOutgoingHull;
    public event EventHandler<IncomingShieldEvent>? OnIncomingShield;
    public event EventHandler<OutgoingShieldEvent>? OnOutgoingShield;
    public event EventHandler<IncomingArmorEvent>? OnIncomingArmor;
    public event EventHandler<OutgoingArmorEvent>? OnOutgoingArmor;
    public event EventHandler<IncomingCapacitorEvent>? OnIncomingCapacitor;
    public event EventHandler<OutgoingCapacitorEvent>? OnOutgoingCapacitor;
    public event EventHandler<IncomingNeutEvent>? OnIncomingNeut;
    public event EventHandler<OutgoingNeutEvent>? OnOutgoingNeut;
    public event EventHandler<IncomingNosEvent>? OnIncomingNos;
    public event EventHandler<OutgoingNosEvent>? OnOutgoingNos;

    private void UpdateRegex(EveLogParserOptions options)
    {
        if (string.IsNullOrWhiteSpace(options.OverviewPath))
        {
            _overviewRegex = null;
            return;
        }
        
        using var reader = new StreamReader(options.OverviewPath);
        var overviewYaml = reader.ReadToEnd();
        
        var deserializer = new Deserializer();
        var settings = deserializer.Deserialize<Dictionary<string, object>>(overviewYaml);
        
        try
        {
            var shipLabels = ExtractShipLabelSettings(settings);
            var shipLabelRegex = BuildOverviewRegex(shipLabels);

            _overviewRegex = shipLabelRegex;
        }
        catch (Exception ex)
        {
            throw new YamlException("Error parsing overview", ex);
        }
    }

    public void StartWatchingCharacter(string characterId)
    {
        _watchedCharacters.Add(characterId);
    }

    public void StopWatchingCharacter(string characterId)
    {
        _watchedCharacters.Remove(characterId);
    }

    private IOrderedEnumerable<IEnumerable<KeyValuePair<string, string?>>> ExtractShipLabelSettings(
        IReadOnlyDictionary<string, object> parsedYaml)
    {
        var shipLabelOrderList = parsedYaml["shipLabelOrder"] as List<object> ??
                                 throw new Exception("shipLabelOrder is not in the overview file");

        var shipLabelOrder = shipLabelOrderList
            .Select(x => x as string ?? "default")
            .ToList();

        var shipLabelList = parsedYaml["shipLabels"] as List<object> ??
                            throw new Exception("shipLabels is not in the overview file");

        return shipLabelList
            .Select(x =>
            {
                // Need to cast into an iterable list
                var itemList = x as List<object>
                               ?? throw new Exception("Unexpected value within the shipLabels");

                return itemList
                    // Each item in the list is a dictionary, cant cast as List<List<object>> for some reason
                    .Select(y => y as List<object>)!
                    // Compile our list of lists into key value pairs
                    .SelectMany<List<object>, KeyValuePair<string, string?>>(w =>
                    {
                        // The first key in each of these is not a list, so it casts to null, and we skip it
                        if (w == null) return new List<KeyValuePair<string, string?>>();

                        return w.Select(z =>
                        {
                            var zList = z as List<object>
                                        ?? throw new Exception("Unexpected value within the shipLabels");


                            var label = zList[0] as string
                                        ?? throw new Exception("Unexpected value within the shipLabels");

                            var value = zList[1] as string;
                            return new KeyValuePair<string, string?>(label, value);
                        });
                    });
            })
            .OrderBy(x =>
            {
                var typePair = x.FirstOrDefault(y => y.Key == "type");
                return shipLabelOrder.IndexOf(typePair.Value ?? "default");
            });
    }

    private string BuildOverviewRegex(IOrderedEnumerable<IEnumerable<KeyValuePair<string, string?>>> labelSettings)
    {
        var shipLabelRegex = new StringBuilder("(?:(?:.*ffffffff>");
        foreach (var labelSetting in labelSettings)
        {
            var label = labelSetting.ToList();

            var safePre = Regex.Escape(label.FirstOrDefault(y => y.Key == "pre").Value ?? string.Empty);
            var safePost = Regex.Escape(label.FirstOrDefault(y => y.Key == "post").Value ?? string.Empty);

            var type = label.FirstOrDefault(x => x.Key == "type").Value;
            var state = label.FirstOrDefault(y => y.Key == "state").Value;
            if (string.IsNullOrWhiteSpace(state) && type is "pilot name")
                shipLabelRegex.Append("(?<Name>)");
            else if (string.IsNullOrWhiteSpace(state) && type is "ship type") shipLabelRegex.Append("(?<Ship>)");

            switch (type)
            {
                case "default":
                {
                    shipLabelRegex.Append((string?)$"(?:{safePre})?");
                    break;
                }
                case "alliance":
                {
                    shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<Alliance>.*?){safePost})?");
                    break;
                }
                case "corporation":
                {
                    shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<Corporation>.*?){safePost})?");
                    break;
                }
                case "ship name":
                {
                    shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<ShipName>.*?){safePost})?");
                    break;
                }
                case "pilot name":
                {
                    shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<Pilot>.*?){safePost})");
                    break;
                }
                case "ship type":
                {
                    shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<Ship>.*?){safePost})");
                    break;
                }
            }
        }

        shipLabelRegex.Append(".*> \\-(?: (?:<Localized .*?>)?(?<Weapon>.*?)(?: \\-|<)|.*))");
        shipLabelRegex.Append($"|{EnglishRegex.PilotAndWeapon})?");
        return shipLabelRegex.ToString();
    }

    private void HandleLogFileRead(object source, LogFileReadEventArgs e)
    {
        Log.Debug("Processing log entry");
        if (!_watchedCharacters.Contains(e.CharacterId)) return;

        var lines = Encoding.UTF8.GetString(e.Content)
            .Split(new[] { "\r\n", "\r", "\n" }, StringSplitOptions.None)
            .ToList();

        foreach (var line in lines)
        {
            CallUntilTrueReturned(e.CharacterId, line,
                FindAndRaiseIncomingDamage,
                FindAndRaiseOutgoingDamage,
                FindAndRaiseIncomingShield,
                FindAndRaiseOutgoingShield,
                FindAndRaiseIncomingArmor,
                FindAndRaiseOutgoingArmor,
                FindAndRaiseIncomingCapacitor,
                FindAndRaiseOutgoingCapacitor,
                FindAndRaiseIncomingNeut,
                FindAndRaiseOutgoingNeut,
                FindAndRaiseIncomingNos,
                FindAndRaiseOutgoingNos,
                FindAndRaiseIncomingHull,
                FindAndRaiseOutgoingHull);
        }
    }

    /// <summary>
    ///     Calls each function with the given params until one function returns true
    ///     This allows us to quit early if we find a match, without a gigantic if chain
    /// </summary>
    /// <param name="characterId"></param>
    /// <param name="logLine"></param>
    /// <param name="functions"></param>
    private static void CallUntilTrueReturned(string characterId, string logLine,
        params Func<string, string, bool>[] functions)
    {
        foreach (var func in functions)
            if (func(characterId, logLine))
                break;
    }

    private bool FindAndRaiseEvent<T>(EventHandler<T>? eventHandler, string constantRegex,
        Func<DateTimeOffset, string, int, string, string, string, string, string, string, T> argsFactory, string characterId,
        string logLine, bool useOverviewRegex = true)
        where T : EveLogEventArgs
    {
        if (eventHandler == null) return false;

        var charRegex = useOverviewRegex && _overviewRegex != null
            ? _overviewRegex
            : EnglishRegex.PilotAndWeapon;

        var regex = new Regex(EnglishRegex.Timestamp + constantRegex + charRegex);
        var match = regex.Match(logLine);
        if (!match.Success) return false;
        Log.Debug("Log matches type {EventType}", typeof(T).Name);

        
        var timestampStr = match.Groups.GetValueOrDefault("Timestamp")?.Value ?? "Unknown";
        if (!DateTime.TryParseExact(timestampStr, "yyyy.MM.dd HH:mm:ss", null,
                DateTimeStyles.AssumeUniversal, out var timestamp))
        {
            timestamp = DateTime.UtcNow;
        }
        
        
        var amountReceived = int.Parse(match.Groups.GetValueOrDefault("Amount")?.Value ?? "0");
        var toName = match.Groups.GetValueOrDefault("Pilot")?.Value ?? "Unknown";
        var toShip = match.Groups.GetValueOrDefault("Ship")?.Value ?? "Unknown";
        var weapon = match.Groups.GetValueOrDefault("Weapon")?.Value ?? "Unknown";
        var application = match.Groups.GetValueOrDefault("Application")?.Value ?? "Unknown";
        var corporation = match.Groups.GetValueOrDefault("Corporation")?.Value ?? "Unknown";
        var alliance = match.Groups.GetValueOrDefault("Alliance")?.Value ?? "Unknown";
        var newEvent = argsFactory(timestamp, characterId, amountReceived, toName, toShip, weapon, application, corporation,alliance);

        eventHandler(this, newEvent);
        return true;
    }

    private bool FindAndRaiseIncomingNos(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new IncomingNosEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnIncomingNos, EnglishRegex.IncomingNos,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseOutgoingNos(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingNosEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingNos, EnglishRegex.OutgoingNos,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseIncomingNeut(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new IncomingNeutEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnIncomingNeut, EnglishRegex.IncomingNeut,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseOutgoingNeut(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingNeutEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingNeut, EnglishRegex.OutgoingNeut,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseIncomingShield(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new IncomingShieldEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnIncomingShield, EnglishRegex.IncomingShield,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseOutgoingShield(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingShieldEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingShield, EnglishRegex.OutgoingShield,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseIncomingArmor(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new IncomingArmorEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnIncomingArmor, EnglishRegex.IncomingArmor,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseOutgoingArmor(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingArmorEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingArmor, EnglishRegex.OutgoingArmor,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseIncomingCapacitor(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new IncomingCapacitorEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnIncomingCapacitor, EnglishRegex.IncomingCapacitor,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseOutgoingCapacitor(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingCapacitorEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingCapacitor, EnglishRegex.OutgoingCapacitor,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseIncomingHull(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new IncomingHullEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnIncomingHull, EnglishRegex.IncomingHull,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseOutgoingHull(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingHullEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingHull, EnglishRegex.OutgoingHull,
            argsBuilder, characterId, logLine);
    }

    private bool FindAndRaiseIncomingDamage(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new IncomingDamageEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnIncomingDamage, EnglishRegex.IncomingDamage,
            argsBuilder, characterId, logLine, false);
    }

    private bool FindAndRaiseOutgoingDamage(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingDamageEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingDamage, EnglishRegex.OutgoingDamage,
            argsBuilder, characterId, logLine, false);
    }

    public void Dispose()
    {
        _logReaderService.OnFileRead -= HandleLogFileRead;
    }
}