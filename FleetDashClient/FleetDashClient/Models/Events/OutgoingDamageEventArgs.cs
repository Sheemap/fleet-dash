using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events
{
	public class OutgoingDamageEventArgs : EveLogEvent
	{
		public OutgoingDamageEventArgs(string characterId, int amount, string pilot, string ship, string weapon, string application) : base(characterId, amount, pilot, ship, weapon)
		{
			Application = application;
		}
		
		public string Application { get; }

	}
}
