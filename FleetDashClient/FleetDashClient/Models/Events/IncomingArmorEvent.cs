﻿namespace FleetDashClient.Models.Events;

public class IncomingArmorEvent : EveLogEventArgs
{
    public IncomingArmorEvent(DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship,
        string weapon, string application, string corporation, string alliance) :
        base(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance)
    {
    }
}