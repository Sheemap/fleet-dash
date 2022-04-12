using System.Globalization;
using System.Text;
using System.Text.RegularExpressions;
using EveLogParser.Constants;
using EveLogParser.Models;
using EveLogParser.Models.Events;
using Serilog;

namespace EveLogParser;

public class LogParserService : ILogParserService, IDisposable
{
    private readonly ILogReaderService _logReaderService;

    private readonly List<WatchedCharacter> _watchedCharacters = new();

    public LogParserService(ILogReaderService logReaderService)
    {
        _logReaderService = logReaderService;
        _logReaderService.OnFileRead += HandleLogFileRead;
    }

    public void Dispose()
    {
        _logReaderService.OnFileRead -= HandleLogFileRead;
    }

    public Task Start()
    {
        return _logReaderService.Start();
    }

    public void Stop()
    {
        _logReaderService.Stop();
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

    public void StartWatchingCharacter(string characterId, string? overviewPath)
    {
        var a = new WatchedCharacter(characterId, overviewPath);
        _watchedCharacters.Add(new WatchedCharacter(characterId, overviewPath));
    }

    public void StopWatchingCharacter(string characterId)
    {
        var toRemove = _watchedCharacters.FirstOrDefault(x => x.CharacterId == characterId);
        if (toRemove != null) _watchedCharacters.Remove(toRemove);
    }

    private bool IsCharacterWatched(string characterId)
    {
        return _watchedCharacters.Any(x => x.CharacterId == characterId);
    }

    private void HandleLogFileRead(object? source, LogFileReadEventArgs e)
    {
        Log.Debug("Processing log entry");
        if (!IsCharacterWatched(e.CharacterId))
        {
            Log.Debug("No character matched for log, not parsing.");
            return;
        }

        var lines = Encoding.UTF8.GetString(e.Content)
            .Split(new[] {"\r\n", "\r", "\n"}, StringSplitOptions.None)
            .ToList();
        
        Log.Debug("Processing {LineCount} log lines.", lines.Count);

        var watchedCharacter = _watchedCharacters.First(x => x.CharacterId == e.CharacterId);

        foreach (var line in lines)
            CallUntilTrueReturned(watchedCharacter, line,
                BlockNoParse,
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
        
        Log.Debug("Finished processing log entry");
    }

    /// <summary>
    ///     Calls each function with the given params until one function returns true
    ///     This allows us to quit early if we find a match, without a gigantic if chain
    /// </summary>
    /// <param name="watchedCharacter"></param>
    /// <param name="logLine"></param>
    /// <param name="functions"></param>
    private static void CallUntilTrueReturned(WatchedCharacter watchedCharacter, string logLine,
        params Func<WatchedCharacter, string, bool>[] functions)
    {
        foreach (var func in functions)
            if (func(watchedCharacter, logLine))
                break;
    }

    private bool FindAndRaiseEvent<T>(EventHandler<T>? eventHandler, string constantRegex,
        Func<DateTimeOffset, string, int, string, string, string, string, string, string, T> argsFactory,
        WatchedCharacter watchedCharacter,
        string logLine, bool useDefault = false)
        where T : EveLogEventArgs
    {
        if (eventHandler == null) return false;

        var regexString = EnglishRegex.Timestamp + constantRegex;

        var regex = new Regex(regexString);
        var match = regex.Match(logLine);
        if (!match.Success) return false;
        Log.Debug("Log matches type {EventType}", typeof(T).Name);

        var parsedLabels = ParseShipLabels(watchedCharacter, logLine);
        if (parsedLabels == null) return false;
        
        var allMatches = new Dictionary<string, string>();
        foreach (Group parsedLabel in parsedLabels)
        {
            allMatches.TryAdd(parsedLabel.Name, parsedLabel.Value);
        }
        foreach (Group matchGroup in match.Groups)
        {
            allMatches.TryAdd(matchGroup.Name, match.Value);
        }

        ConstructAndRaiseEvent(allMatches, watchedCharacter.CharacterId, argsFactory, eventHandler);
        return true;
    }

    private void ConstructAndRaiseEvent<T>(IReadOnlyDictionary<string, string> matches,
        string characterId,
        Func<DateTimeOffset, string, int, string, string, string, string, string, string, T> argsFactory,
        EventHandler<T> eventHandler)
        where T : EveLogEventArgs
    {
        var timestampStr = matches.GetValueOrDefault("timestamp", "Unknown");
        var timestamp = ParseTimestamp(timestampStr);

        var amountStr = matches.GetValueOrDefault("amount", "0");
        if (!int.TryParse(amountStr, out var amount))
        {
            amount = 0;
        }

        var newEvent = argsFactory(timestamp,
            characterId,
            amount,
            matches.GetValueOrDefault("pilotname", "Unknown"),
            matches.GetValueOrDefault("shiptype", "Unknown"),
            matches.GetValueOrDefault("weapon", "Unknown"),
            matches.GetValueOrDefault("application", "Unknown"),
            matches.GetValueOrDefault("corporation", "Unknown"),
            matches.GetValueOrDefault("alliance", "Unknown"));

        eventHandler(this, newEvent);
    }

    private static GroupCollection? ParseShipLabels(WatchedCharacter character, string logLine)
    {
        // Attempt with Alliance included
        var characterRegex = character.GetCharacterRegex(true);
        var withAllianceRegex = new Regex(characterRegex);
        var shipLabelMatch = withAllianceRegex.Match(logLine);
        if (shipLabelMatch.Success) return shipLabelMatch.Groups;
        
        // Attempt without Alliance
        var characterRegexWithoutAlliance = character.GetCharacterRegex(false);
        var withoutAllianceRegex = new Regex(characterRegexWithoutAlliance);
        var matchWithoutAlliance = withoutAllianceRegex.Match(logLine);
        if (matchWithoutAlliance.Success) return matchWithoutAlliance.Groups;
        
        // We failed :(
        Log.Warning("Failed to parse a log line! Could not extract the ship label values.");
        return null;
    }

    private static DateTimeOffset ParseTimestamp(string timestampStr)
    {
        return !DateTimeOffset.TryParseExact(timestampStr, "yyyy.MM.dd HH:mm:ss", null,
            DateTimeStyles.AssumeUniversal, out var timestamp)
            ? DateTimeOffset.UtcNow
            : timestamp;
    }

    private bool FindAndRaiseIncomingNos(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new IncomingNosEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnIncomingNos, EnglishRegex.IncomingNos,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseOutgoingNos(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new OutgoingNosEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnOutgoingNos, EnglishRegex.OutgoingNos,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseIncomingNeut(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new IncomingNeutEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnIncomingNeut, EnglishRegex.IncomingNeut,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseOutgoingNeut(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new OutgoingNeutEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnOutgoingNeut, EnglishRegex.OutgoingNeut,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseIncomingShield(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new IncomingShieldEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnIncomingShield, EnglishRegex.IncomingShield,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseOutgoingShield(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new OutgoingShieldEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnOutgoingShield, EnglishRegex.OutgoingShield,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseIncomingArmor(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new IncomingArmorEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnIncomingArmor, EnglishRegex.IncomingArmor,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseOutgoingArmor(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new OutgoingArmorEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnOutgoingArmor, EnglishRegex.OutgoingArmor,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseIncomingCapacitor(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new IncomingCapacitorEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnIncomingCapacitor, EnglishRegex.IncomingCapacitor,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseOutgoingCapacitor(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new OutgoingCapacitorEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnOutgoingCapacitor, EnglishRegex.OutgoingCapacitor,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseIncomingHull(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new IncomingHullEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnIncomingHull, EnglishRegex.IncomingHull,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseOutgoingHull(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new OutgoingHullEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnOutgoingHull, EnglishRegex.OutgoingHull,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseIncomingDamage(WatchedCharacter watchedCharacter, string logLine)
    {
        var eventHandler = OnIncomingDamage;
        if (eventHandler == null) return false;
        
        var regex = new Regex(EnglishRegex.IncomingDamage);
        var match = regex.Match(logLine);
        if (!match.Success) return false;

        var matches = new Dictionary<string, string>();
        foreach (Group matchGroup in match.Groups)
        {
            matches.TryAdd(matchGroup.Name, matchGroup.Value);
        }
        
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new IncomingDamageEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        ConstructAndRaiseEvent(matches, watchedCharacter.CharacterId, argsBuilder, eventHandler);
        return true;
    }

    private bool FindAndRaiseOutgoingDamage(WatchedCharacter watchedCharacter, string logLine)
    {
        var eventHandler = OnOutgoingDamage;
        if (eventHandler == null) return false;
        
        var regex = new Regex(EnglishRegex.OutgoingDamage);
        var match = regex.Match(logLine);
        if (!match.Success) return false;

        var matches = new Dictionary<string, string>();
        foreach (Group matchGroup in match.Groups)
        {
            matches.TryAdd(matchGroup.Name, matchGroup.Value);
        }
        
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new OutgoingDamageEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        ConstructAndRaiseEvent(matches, watchedCharacter.CharacterId, argsBuilder, eventHandler);
        return true;
    }

    private bool FindAndRaiseIncomingJam(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new IncomingJamEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnIncomingJam, EnglishRegex.IncomingJam,
            argsBuilder, watchedCharacter, logLine);
    }

    private bool FindAndRaiseOutgoingJam(WatchedCharacter watchedCharacter, string logLine)
    {
        var argsBuilder = (DateTimeOffset timestamp, string charId, int amount, string pilot, string ship,
                string weapon, string application, string corporation, string alliance) =>
            new OutgoingJamEvent(timestamp, charId, amount, pilot, ship, weapon, application, corporation,
                alliance);

        return FindAndRaiseEvent(OnOutgoingJam, EnglishRegex.OutgoingJam,
            argsBuilder, watchedCharacter, logLine);
    }
    
    private static bool BlockNoParse(WatchedCharacter watchedCharacter, string logLine)
    {
        var regex = new Regex(EnglishRegex.NoParse);
        return regex.IsMatch(logLine);
    }
}