using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components.Web;

namespace FleetDashClient.Services
{
	public class LogParserService : ILogParserService
	{
        public event EventHandler<IncomingDamageEventArgs> RaiseIncomingDamageEvent;
        public event EventHandler<OutgoingDamageEventArgs> RaiseOutgoingDamageEvent;
        public event EventHandler<IncomingHullEventArgs> RaiseIncomingHullEvent;
        public event EventHandler<OutgoingHullEventArgs> RaiseOutgoingHullEvent;
        public event EventHandler<IncomingArmorEventArgs> RaiseIncomingArmorEvent;
        public event EventHandler<OutgoingArmorEventArgs> RaiseOutgoingArmorEvent;
        public event EventHandler<IncomingCapacitorEventArgs> RaiseIncomingCapacitorEvent;
        public event EventHandler<OutgoingCapacitorEventArgs> RaiseOutgoingCapacitorEvent;

        
        private readonly Dictionary<string, string?> _watchedCharacters = new();

        public LogParserService(ILogReaderService logReaderService)
        {
            logReaderService.RaiseFileReadEvent += HandleLogFileReadEvent;
        }

        public void StartWatchingCharacter(string characterId, string overviewSettings)
        {
            if (string.IsNullOrWhiteSpace(overviewSettings))
            {
                _watchedCharacters[characterId] = null;
                return;
            }
            
            // Parse overviewSettings as YAML
            var deserializer = new YamlDotNet.Serialization.Deserializer();
            var settings = deserializer.Deserialize<Dictionary<string, object>>(overviewSettings);

            var shipLabelOrder = (settings["shipLabelOrder"] as List<object>)
                .Select(x => x as string ?? "default")
                .ToList();

            var shipLabels = (settings["shipLabels"] as List<object>)
                .Select(x =>
                {
                    // Need to cast into an iterable list
                    var itemList = x as List<object>;

                    return itemList
                        // Each item in the list is a dictionary, cant cast as List<List<object>> for some reason
                        .Select(y => y as List<object>)
                        // Compile our list of lists into key value pairs
                        .SelectMany<List<object>, KeyValuePair<string, string>>(w =>
                        {
                            // The first key in each of these is not a list, so it casts to null, and we skip it
                            if (w == null) return new List<KeyValuePair<string, string>>();

                            return w.Select(z =>
                            {
                                var zList = z as List<object>;
                                var label = zList[0] as string;
                                var value = zList[1] as string;
                                return new KeyValuePair<string, string>(label, value);
                            });
                        });
                })
                .OrderBy(x =>
                {
                    var typePair = x.FirstOrDefault(y => y.Key == "type");
                    return shipLabelOrder.IndexOf(typePair.Value ?? "default");
                });

            var shipLabelRegex = new StringBuilder("(?:(?:.*ffffffff>");
            foreach (var label in shipLabels)
            {
                var type = label.FirstOrDefault(x => x.Key == "type").Value;
                var state = label.FirstOrDefault(y => y.Key == "state").Value;
                if (string.IsNullOrWhiteSpace(state) && type is "pilot name")
                {
                    shipLabelRegex.Append($"(?<Name>)");
                }
                else if (string.IsNullOrWhiteSpace(state) && type is "ship type")
                {
                    shipLabelRegex.Append($"(?<Ship>)");
                }
                
                switch (type)
                {
                    case "default":
                    {
                        var safePre = Regex.Escape(label.FirstOrDefault(y => y.Key == "pre").Value ?? "");
                        shipLabelRegex.Append((string?)$"(?:{safePre})?");
                        break;
                    }
                    case "alliance":
                    {
                        var safePre = Regex.Escape(label.FirstOrDefault(y => y.Key == "pre").Value ?? "");
                        var safePost = Regex.Escape(label.FirstOrDefault(y => y.Key == "post").Value ?? "");
                        shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<Alliance>.*?){safePost})?");
                        break;
                    }
                    case "corporation":
                    {
                        var safePre = Regex.Escape(label.FirstOrDefault(y => y.Key == "pre").Value ?? "");
                        var safePost = Regex.Escape(label.FirstOrDefault(y => y.Key == "post").Value ?? "");
                        shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<Corporation>.*?){safePost})?");
                        break;
                    }
                    case "ship name":
                    {
                        var safePre = Regex.Escape(label.FirstOrDefault(y => y.Key == "pre").Value ?? "");
                        var safePost = Regex.Escape(label.FirstOrDefault(y => y.Key == "post").Value ?? "");
                        shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<ShipName>.*?){safePost})?");
                        break;
                    }
                    case "pilot name":
                    {
                        var safePre = Regex.Escape(label.FirstOrDefault(y => y.Key == "pre").Value ?? "");
                        var safePost = Regex.Escape(label.FirstOrDefault(y => y.Key == "post").Value ?? "");
                        shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<Name>.*?){safePost})");
                        break;
                    }
                    case "ship type":
                    {
                        var safePre = Regex.Escape(label.FirstOrDefault(y => y.Key == "pre").Value ?? "");
                        var safePost = Regex.Escape(label.FirstOrDefault(y => y.Key == "post").Value ?? "");
                        shipLabelRegex.Append((string?)$"(?:{safePre}(?:<localized .*?>)?(?<Ship>.*?){safePost})");
                        break;
                    }
                }
            }
            
            shipLabelRegex.Append(".*> \\-(?: (?:<Localized .*?>)?(?<Weapon>.*?)(?: \\-|<)|.*))");
            shipLabelRegex.Append($"|{Constants.EnglishRegex.PilotAndWeapon})?");

            _watchedCharacters[characterId] = shipLabelRegex.ToString();
        }
        
        public void StopWatchingCharacter(string characterId)
        {
            _watchedCharacters.Remove(characterId);
        }

        private void HandleLogFileReadEvent(object source, LogFileReadEventArgs e)
        {
            if (!_watchedCharacters.ContainsKey(e.CharacterId)) return;
            
            var line = Encoding.UTF8.GetString(e.Content);

            CallUntilTrueReturned(e.CharacterId, line,
                FindAndRaiseIncomingDamage,
                FindAndRaiseOutgoingDamage,
                FindAndRaiseIncomingArmor,
                FindAndRaiseOutgoingArmor,
                FindAndRaiseIncomingCapacitor,
                FindAndRaiseOutgoingCapacitor,
                FindAndRaiseIncomingHull,
                FindAndRaiseOutgoingHull);
        }
        
        /// <summary>
        /// Calls each function with the given params until one function returns true
        /// This allows us to quit early if we find a match, without a gigantic if chain
        /// </summary>
        /// <param name="characterId"></param>
        /// <param name="logLine"></param>
        /// <param name="functions"></param>
        private static void CallUntilTrueReturned(string characterId, string logLine,
            params Func<string, string, bool> [] functions)
        {
            foreach (var func in functions)
            {
                if (func(characterId, logLine)) break;
            }
        }
        
        private bool FindAndRaiseEvent<T>(EventHandler<T> eventHandler, string constantRegex,
            Func<string, int, string, string,string, T> argsFactory, string characterId, string logLine)
            where T : EveLogEvent
        {
            if (eventHandler == null) return false;
        
            var charRegex = _watchedCharacters.GetValueOrDefault(characterId, Constants.EnglishRegex.PilotAndWeapon);
        
            var regex = new Regex(constantRegex + charRegex);
            var match = regex.Match(logLine);
            if (!match.Success) return false;
            
            var amountReceived = Int32.Parse(match.Groups.GetValueOrDefault("Amount")?.Value ?? "0");
            var toName = match.Groups.GetValueOrDefault("Name")?.Value ?? "Unknown";
            var toShip = match.Groups.GetValueOrDefault("Ship")?.Value ?? "Unknown";
            var weapon = match.Groups.GetValueOrDefault("Weapon")?.Value ?? "Unknown";
            var newEvent = argsFactory(characterId, amountReceived, toName, toShip, weapon);
                
            eventHandler(this, newEvent);
            return true;
        }

        private bool FindAndRaiseIncomingArmor(string characterId, string logLine)
        {
            var argsBuilder = (string characterId, int amount, string pilot, string ship, string weapon) =>
                new IncomingArmorEventArgs(characterId, amount, pilot, ship, weapon);
            
            return FindAndRaiseEvent(RaiseIncomingArmorEvent, Constants.EnglishRegex.IncomingArmor,
                argsBuilder, characterId, logLine);
        }
        
        private bool FindAndRaiseOutgoingArmor(string characterId, string logLine)
        {
            var argsBuilder = (string characterId, int amount, string pilot, string ship, string weapon) =>
                new OutgoingArmorEventArgs(characterId, amount, pilot, ship, weapon);
            
            return FindAndRaiseEvent(RaiseOutgoingArmorEvent, Constants.EnglishRegex.OutgoingArmor,
                argsBuilder, characterId, logLine);
        }
        
        private bool FindAndRaiseIncomingCapacitor(string characterId, string logLine)
        {
            var argsBuilder = (string characterId, int amount, string pilot, string ship, string weapon) =>
                new IncomingCapacitorEventArgs(characterId, amount, pilot, ship, weapon);
            
            return FindAndRaiseEvent(RaiseIncomingCapacitorEvent, Constants.EnglishRegex.IncomingCapacitor,
                argsBuilder, characterId, logLine);
        }
        
        private bool FindAndRaiseOutgoingCapacitor(string characterId, string logLine)
        {
            var argsBuilder = (string characterId, int amount, string pilot, string ship, string weapon) =>
                new OutgoingCapacitorEventArgs(characterId, amount, pilot, ship, weapon);
            
            return FindAndRaiseEvent(RaiseOutgoingCapacitorEvent, Constants.EnglishRegex.OutgoingCapacitor,
                argsBuilder, characterId, logLine);
        }
        
        private bool FindAndRaiseIncomingHull(string characterId, string logLine)
        {
            var argsBuilder = (string characterId, int amount, string pilot, string ship, string weapon) =>
                new IncomingHullEventArgs(characterId, amount, pilot, ship, weapon);
            
            return FindAndRaiseEvent(RaiseIncomingHullEvent, Constants.EnglishRegex.IncomingHull,
                argsBuilder, characterId, logLine);
        }
        
        private bool FindAndRaiseOutgoingHull(string characterId, string logLine)
        {
            var argsBuilder = (string characterId, int amount, string pilot, string ship, string weapon) =>
                new OutgoingHullEventArgs(characterId, amount, pilot, ship, weapon);
            
            return FindAndRaiseEvent(RaiseOutgoingHullEvent, Constants.EnglishRegex.OutgoingHull,
                argsBuilder, characterId, logLine);
        }
        
        private bool FindAndRaiseIncomingDamage(string characterId, string logLine)
        {
            var raiseDamage = RaiseIncomingDamageEvent;
            if (raiseDamage == null) return false;

            var regex = new Regex(Constants.EnglishRegex.IncomingDamage + Constants.EnglishRegex.PilotAndWeapon);
            var match = regex.Match(logLine);
            if (!match.Success) return false;
            
            var amountReceived = Int32.Parse(match.Groups.GetValueOrDefault("Amount")?.Value ?? "0");
            var fromName = match.Groups.GetValueOrDefault("Name")?.Value ?? "Unknown";
            var fromShip = match.Groups.GetValueOrDefault("Ship")?.Value ?? "Unknown";
            var weapon = match.Groups.GetValueOrDefault("Weapon")?.Value ?? "Unknown";
            var application = match.Groups.GetValueOrDefault("Application")?.Value ?? "Unknown";
            var newEvent = new IncomingDamageEventArgs(characterId, amountReceived, fromName, fromShip,
                weapon, application);

            raiseDamage(this, newEvent);
            return true;

        }

        private bool FindAndRaiseOutgoingDamage(string characterId, string logLine)
        {
            var raiseDamage = RaiseOutgoingDamageEvent;
            if (raiseDamage == null) return false;
            
            var regex = new Regex(Constants.EnglishRegex.OutgoingDamage + Constants.EnglishRegex.PilotAndWeapon);
            var match = regex.Match(logLine);
            if (!match.Success) return false;
            
            var amount = Int32.Parse(match.Groups.GetValueOrDefault("Amount")?.Value ?? "0");
            var toName = match.Groups.GetValueOrDefault("Name")?.Value ?? "Unknown";
            var toShip = match.Groups.GetValueOrDefault("Ship")?.Value ?? "Unknown";
            var weapon = match.Groups.GetValueOrDefault("Weapon")?.Value ?? "Unknown";
            var application = match.Groups.GetValueOrDefault("Application")?.Value ?? "Unknown";
            var newEvent = new OutgoingDamageEventArgs(characterId, amount, toName, toShip, weapon,
                application);

            raiseDamage(this, newEvent);
            return true;

        }
    }
}
