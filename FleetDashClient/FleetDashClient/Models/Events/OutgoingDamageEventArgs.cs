using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events
{
	public class OutgoingDamageEventArgs : EventArgs
	{
		public OutgoingDamageEventArgs(string characterId, int amount, string toName, string toTag, string toShip, string weapon, string application)
		{
			CharacterId = characterId;
			Amount = amount;
			ToName = toName;
			ToTag = toTag;
			ToShip = toShip;
			Weapon = weapon;
			Application = application;
		}
		
		public string CharacterId { get; }
		public int Amount { get; }
		public string ToName { get; }
		public string ToTag { get; }
		public string ToShip { get; }
		public string Weapon { get; }
		public string Application { get; }
	}
}
