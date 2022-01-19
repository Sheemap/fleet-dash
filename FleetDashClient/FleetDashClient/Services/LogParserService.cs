using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace FleetDashClient.Services
{
	public class LogParserService : ILogParserService
	{
        public event EventHandler<IncomingDamageEventArgs> RaiseIncomingDamageEvent;
        public event EventHandler<OutgoingDamageEventArgs> RaiseOutgoingDamageEvent;
        public event EventHandler<OutgoingArmorEventArgs> RaiseOutgoingArmorEvent;
        
        private readonly Dictionary<string, string?> _watchedCharacters = new Dictionary<string, string?>();

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
            FindAndRaiseIncomingDamage(e.CharacterId, line);
            FindAndRaiseOutgoingDamage(e.CharacterId, line);
            FindAndRaiseOutgoingArmor(e.CharacterId, line);
        }

        private void FindAndRaiseIncomingArmor(string characterId, string logLine)
        {
            
        }

        private void FindAndRaiseOutgoingArmor(string characterId, string logLine)
        {
            var raiseArmor = RaiseOutgoingArmorEvent;
            if (raiseArmor == null) return;

            var charRegex = _watchedCharacters.GetValueOrDefault(characterId, Constants.EnglishRegex.PilotAndWeapon);

            var regex = new Regex(Constants.EnglishRegex.OutgoingArmor + charRegex);
            var match = regex.Match(logLine);
            if (match.Success && match.Groups.Count >= 6)
            {
                var amountReceived = Int32.Parse(match.Groups.GetValueOrDefault("Amount")?.Value ?? "0");
                var toName = match.Groups.GetValueOrDefault("Name")?.Value ?? "Unknown";
                var toShip = match.Groups.GetValueOrDefault("Ship")?.Value ?? "Unknown";
                var weapon = match.Groups.GetValueOrDefault("Weapon")?.Value ?? "Unknown";
                var newEvent = new OutgoingArmorEventArgs(characterId, amountReceived, toName, toShip, weapon);

                raiseArmor(this, newEvent);
            }
        }

        private void FindAndRaiseIncomingDamage(string characterId, string logLine)
        {
            var raiseDamage = RaiseIncomingDamageEvent;
            if (raiseDamage == null) return;

            var regex = new Regex(Constants.EnglishRegex.IncomingDamage + Constants.EnglishRegex.PilotAndWeapon);
            var match = regex.Match(logLine);
            if (match.Success && match.Groups.Count >= 7)
            {
                var amountReceived = Int32.Parse(match.Groups.GetValueOrDefault("Amount")?.Value ?? "0");
                var fromName = match.Groups.GetValueOrDefault("Name")?.Value ?? "Unknown";
                var fromTag = match.Groups.GetValueOrDefault("Tag")?.Value ?? "Unknown";
                var fromShip = match.Groups.GetValueOrDefault("Ship")?.Value ?? "Unknown";
                var weapon = match.Groups.GetValueOrDefault("Weapon")?.Value ?? "Unknown";
                var application = match.Groups.GetValueOrDefault("Application")?.Value ?? "Unknown";
                var newEvent = new IncomingDamageEventArgs(characterId, amountReceived, fromName, fromTag, fromShip,
                    weapon, application);

                raiseDamage(this, newEvent);
            }
        }

        private void FindAndRaiseOutgoingDamage(string characterId, string logLine)
        {
            var raiseDamage = RaiseOutgoingDamageEvent;
            if (raiseDamage == null) return;
            
            var regex = new Regex(Constants.EnglishRegex.OutgoingDamage + Constants.EnglishRegex.PilotAndWeapon);
            var match = regex.Match(logLine);
            if (match.Success && match.Groups.Count >= 7)
            {
                var amount = Int32.Parse(match.Groups.GetValueOrDefault("Amount")?.Value ?? "0");
                var toName = match.Groups.GetValueOrDefault("Name")?.Value ?? "Unknown";
                var toTag = match.Groups.GetValueOrDefault("Tag")?.Value ?? "Unknown";
                var toShip = match.Groups.GetValueOrDefault("Ship")?.Value ?? "Unknown";
                var weapon = match.Groups.GetValueOrDefault("Weapon")?.Value ?? "Unknown";
                var application = match.Groups.GetValueOrDefault("Application")?.Value ?? "Unknown";
                var newEvent = new OutgoingDamageEventArgs(characterId, amount, toName, toTag, toShip, weapon,
                    application);

                raiseDamage(this, newEvent);
            }
        }
    }
}
