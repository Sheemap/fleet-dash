using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events
{
	public class OutgoingHullEventArgs : EveLogEvent
	{
		public OutgoingHullEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(characterId, amount, pilot, ship, weapon)
		{
		}
	}
}
