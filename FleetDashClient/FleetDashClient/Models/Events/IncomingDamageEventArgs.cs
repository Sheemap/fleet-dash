using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events
{
	public class IncomingDamageEventArgs : EventArgs
	{
		public IncomingDamageEventArgs(string characterID, int amount, string fromName, string fromTag, string ship, string weapon, string application)
		{
			CharacterID = characterID;
			Amount = amount;
			FromName = fromName;
			FromTag = fromTag;
			Ship = ship;
			Weapon = weapon;
			Application = application;
		}
		
		public string CharacterID { get; set; }
		public int Amount { get; set; }
		public string FromName { get; set; }
		public string FromTag { get; set; }
		public string Ship { get; set; }
		public string Weapon { get; set; }
		public string Application { get; set; }
	}
}
