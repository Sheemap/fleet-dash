using System.Collections.Immutable;
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
    private readonly ImmutableList<string> _shipLabelPriority = new List<string>
    {
        "ship type",
        "pilot name",
        "ship name",
        "default",
        "corporation",
        "alliance",
    }.ToImmutableList();
    
    private readonly ImmutableList<string> _defaultShipLabelOrder = new List<string>
    {
        "pilot name",
        "corporation",
        "ship type",
        "default",
    }.ToImmutableList();
    
    private readonly List<string> _watchedCharacters = new();
    private ImmutableList<string> _shipLabelOrder;
    
    private readonly ILogReaderService _logReaderService;
    
    public Task Start() => _logReaderService.Start();
    public void Stop() => _logReaderService.Stop();

    public LogParserService(ILogReaderService logReaderService, IOptionsMonitor<EveLogParserOptions> options)
    {
        _logReaderService = logReaderService;
        _logReaderService.OnFileRead += HandleLogFileRead;
        
        UpdateShipOrder(options.CurrentValue);
        options.OnChange(UpdateShipOrder);
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
    public event EventHandler<IncomingJamEvent>? OnIncomingJam;
    public event EventHandler<OutgoingJamEvent>? OnOutgoingJam;

    
    private void UpdateShipOrder(EveLogParserOptions options)
    {
        if (string.IsNullOrWhiteSpace(options.OverviewPath))
        {
            _shipLabelOrder = _defaultShipLabelOrder;
            return;
        }
        
        using var reader = new StreamReader(options.OverviewPath);
        var overviewYaml = reader.ReadToEnd();
        
        var deserializer = new Deserializer();
        var settings = deserializer.Deserialize<Dictionary<string, object>>(overviewYaml);
        
        try
        {
            var shipLabelOrder = ExtractShipLabelSettings(settings)
                .Select(x => 
                    x.FirstOrDefault(y => y.Key == "type").Value ?? "default")
                .OrderBy(x => x == "default")
                .ToImmutableList();

            _shipLabelOrder = shipLabelOrder;
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

    // Parse the overview yaml, extract the objects into key value pairs
    // Hard to keep a map in your mind of this, maybe its worth building a specific model?
    private static IEnumerable<IEnumerable<KeyValuePair<string, string?>>> ExtractShipLabelSettings(
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
            .Where(x =>
            {
                // Only include items that have `state == 1` which means they are visible
                if (!int.TryParse(x.FirstOrDefault(y => y.Key == "state").Value, out var val))
                {
                    return false;
                }
                return val == 1;
            })
            .OrderBy(x =>
            {
                var typePair = x.FirstOrDefault(y => y.Key == "type");
                return shipLabelOrder.IndexOf(typePair.Value ?? "default");
            });
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
                FindAndRaiseOutgoingHull,
                FindAndRaiseIncomingJam,
                FindAndRaiseOutgoingJam);
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
        Func<DateTimeOffset, string, int, string, string, string, string, string, string, T> argsFactory,
        string characterId,
        string logLine, bool useDefault = false)
        where T : EveLogEventArgs
    {
        if (eventHandler == null) return false;

        var shipLabelGroupsRegex = useDefault ? EnglishRegex.DefaultShipLabels : EnglishRegex.ShipLabelGroups;
        var regexString = EnglishRegex.Timestamp + constantRegex + shipLabelGroupsRegex;
        
        var regex = new Regex(regexString);
        var match = regex.Match(logLine);
        if (!match.Success) return false;
        Log.Debug("Log matches type {EventType}", typeof(T).Name);
        
        var timestampStr = match.Groups.GetValueOrDefault("Timestamp")?.Value ?? "Unknown";
        var timestamp = ParseTimestamp(timestampStr);
        
        var amountReceived = int.Parse(match.Groups.GetValueOrDefault("Amount")?.Value ?? "0");
        
        var shipLabelGroups = match.Groups.Where<Group>(x => x.Name == "ShipLabel")
            .SelectMany(x => x.Captures)
            .Select(x => CleanMatch(x.Value))
            .Where(x => !string.IsNullOrWhiteSpace(x))
            .ToList();

        var labels = GetLabels(shipLabelGroups.Count, useDefault);

        // Populate dictionary with our values and labels
        var dict = new Dictionary<string, string>();
        for (var i = 0; i < labels.Count; i++)
        {
            var label = labels[i];
            var matchedValue = shipLabelGroups[i];
            
            dict.Add(label, matchedValue);
        }

        // Im not amazing at regex, and was struggling to get weapon and application parsed out separately
        // So am doing it later in code here
        var def = dict.GetValueOrDefault("default", "Unknown").Split(" - ");
        var weapon = def[0];
        var application = def.Length > 1 ? def[1] : "Unknown";
        
        
        var newEvent = argsFactory(timestamp,
            characterId, 
            amountReceived,
            dict.GetValueOrDefault("pilot name", "Unknown"),
            dict.GetValueOrDefault("ship type", "Unknown"),
            weapon,
            application,
            dict.GetValueOrDefault("corporation", "Unknown"),
            dict.GetValueOrDefault("alliance", "Unknown"));

        eventHandler(this, newEvent);
        return true;
    }

    private static string CleanMatch(string matchedText)
    {
        return matchedText.Trim().Trim('-').Trim(' ');
    }

    /// <summary>
    /// Gets the labels for specific count of matched groups
    /// Will take into account overview settings for hidden groups, and priority of those groups
    /// </summary>
    /// <param name="count"></param>
    /// <param name="useDefault"></param>
    /// <returns></returns>
    private ImmutableList<string> GetLabels(int count, bool useDefault)
    {
        var order = useDefault ? _defaultShipLabelOrder : _shipLabelOrder;
        return order.Count == count ?
            order :
            // Filter to only fields in our order list, then take only the amount we need, and order according to the order list
            _shipLabelPriority.Where(x => order.Contains(x)).Take(count).OrderBy(x => order.IndexOf(x)).ToImmutableList();
    }

    private static DateTimeOffset ParseTimestamp(string timestampStr)
    {
        return !DateTimeOffset.TryParseExact(timestampStr, "yyyy.MM.dd HH:mm:ss", null,
            DateTimeStyles.AssumeUniversal, out var timestamp) ? DateTimeOffset.UtcNow : timestamp;
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
            argsBuilder, characterId, logLine, true);
    }

    private bool FindAndRaiseOutgoingDamage(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingDamageEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingDamage, EnglishRegex.OutgoingDamage,
            argsBuilder, characterId, logLine, true);
    }
    
    private bool FindAndRaiseIncomingJam(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new IncomingJamEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnIncomingJam, EnglishRegex.IncomingJam,
            argsBuilder, characterId, logLine);
    }
    
    private bool FindAndRaiseOutgoingJam(string characterId, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon, string application, string corporation, string alliance) =>
            new OutgoingJamEvent(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance);

        return FindAndRaiseEvent(OnOutgoingJam, EnglishRegex.OutgoingJam,
            argsBuilder, characterId, logLine);
    }

    public void Dispose()
    {
        _logReaderService.OnFileRead -= HandleLogFileRead;
    }
}