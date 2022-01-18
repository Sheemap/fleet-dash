﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events
{
	public class IncomingDamageEventArgs : EventArgs
	{
		public IncomingDamageEventArgs(string characterId, int amount, string fromName, string fromTag, string fromShip, string weapon, string application)
		{
			CharacterId = characterId;
			Amount = amount;
			FromName = fromName;
			FromTag = fromTag;
			FromShip = fromShip;
			Weapon = weapon;
			Application = application;
		}
		
		public string CharacterId { get; }
		public int Amount { get; }
		public string FromName { get; }
		public string FromTag { get; }
		public string FromShip { get; }
		public string Weapon { get; }
		public string Application { get; }
	}
}
