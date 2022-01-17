using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Services
{
	public class LogParserService : ILogParserService
	{
        public event EventHandler<IncomingDamageEventArgs> RaiseDamageReceivedEvent;

        private readonly ILogReaderService _logReaderService;

        public LogParserService(ILogReaderService logReaderService)
        {
            _logReaderService = logReaderService;

            _logReaderService.RaiseFileReadEvent += HandleLogFileReadEvent;
        }

        public void HandleLogFileReadEvent(object source, LogFileReadEventArgs e)
        {
            var raiseDamage = RaiseDamageReceivedEvent;
            if (raiseDamage == null) return;
            
            var newEvent = new IncomingDamageEventArgs(e.CharacterID, -1, null, null, null, null, null);
            raiseDamage(this, newEvent);
        }
    }
}
