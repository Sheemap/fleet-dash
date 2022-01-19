using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events
{
	public class IncomingArmorEventArgs : EventArgs
	{
		public IncomingArmorEventArgs(string characterId, int amount, string fromName, string fromShip, string weapon)
		{
			CharacterId = characterId;
			Amount = amount;
			FromName = fromName;
			FromShip = fromShip;
			Weapon = weapon;
		}
		
		public string CharacterId { get; }
		public int Amount { get; }
		public string FromName { get; }
		public string FromShip { get; }
		public string Weapon { get; }
	}
}
