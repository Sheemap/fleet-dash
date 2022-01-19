using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events
{
	public class OutgoingCapacitorEventArgs : EventArgs
	{
		public OutgoingCapacitorEventArgs(string characterId, int amount, string toName, string toShip, string weapon)
		{
			CharacterId = characterId;
			Amount = amount;
			ToName = toName;
			ToShip = toShip;
			Weapon = weapon;
		}
		
		public string CharacterId { get; }
		public int Amount { get; }
		public string ToName { get; }
		public string ToShip { get; }
		public string Weapon { get; }
	}
}
