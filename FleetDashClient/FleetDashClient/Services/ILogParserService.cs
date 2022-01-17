using FleetDashClient.Models;
using FleetDashClient.Models.Events;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Services
{
	public interface ILogParserService
	{
		public event EventHandler<IncomingDamageEventArgs> RaiseDamageReceivedEvent;
	}
}
