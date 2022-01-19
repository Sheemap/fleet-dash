using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events;

public class IncomingArmorEventArgs : EveLogEvent
{
    public IncomingArmorEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(characterId, amount, pilot, ship, weapon)
    {
    }
}

