﻿namespace EveLogParser.Models.Events;

public class OutgoingDamageEvent : EveLogEventArgs
{
    public OutgoingDamageEvent(DateTimeOffset timestamp, string characterId, int amount, string pilot, string ship, string weapon,
        string application, string corporation, string alliance) :
        base(timestamp, characterId, amount, pilot, ship, weapon, application, corporation, alliance)
    {
    }
}