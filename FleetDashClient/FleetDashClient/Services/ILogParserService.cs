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
		public event EventHandler<IncomingDamageEventArgs> RaiseIncomingDamageEvent;
		public event EventHandler<OutgoingDamageEventArgs> RaiseOutgoingDamageEvent;
		void StartWatchingCharacter(string characterId, string overviewSettings);
		void StopWatchingCharacter(string characterId);
		event EventHandler<OutgoingArmorEventArgs> RaiseOutgoingArmorEvent;
		event EventHandler<IncomingArmorEventArgs> RaiseIncomingArmorEvent;
		event EventHandler<IncomingCapacitorEventArgs> RaiseIncomingCapacitorEvent;
		event EventHandler<OutgoingCapacitorEventArgs> RaiseOutgoingCapacitorEvent;
		event EventHandler<IncomingHullEventArgs> RaiseIncomingHullEvent;
		event EventHandler<OutgoingHullEventArgs> RaiseOutgoingHullEvent;
		event EventHandler<IncomingShieldEventArgs> RaiseIncomingShieldEvent;
		event EventHandler<OutgoingShieldEventArgs> RaiseOutgoingShieldEvent;
		event EventHandler<IncomingNeutEventArgs> RaiseIncomingNeutEvent;
		event EventHandler<OutgoingNeutEventArgs> RaiseOutgoingNeutEvent;
		event EventHandler<IncomingNosEventArgs> RaiseIncomingNosEvent;
		event EventHandler<OutgoingNosEventArgs> RaiseOutgoingNosEvent;
	}
}
