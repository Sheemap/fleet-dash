using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using System;
using System.Collections.Generic;
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
        
        private readonly Dictionary<string, Regex?> _watchedCharacters = new Dictionary<string, Regex?>();

        public LogParserService(ILogReaderService logReaderService)
        {
            logReaderService.RaiseFileReadEvent += HandleLogFileReadEvent;
        }
        
        public void StartWatchingCharacter(string characterId, string overviewSettings)
        {
            // TODO: overviewSettings should be valid YAML and be parsed to create a regex
            _watchedCharacters[characterId] = null;
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
