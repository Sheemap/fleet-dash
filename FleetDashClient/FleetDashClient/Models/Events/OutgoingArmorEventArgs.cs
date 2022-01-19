﻿namespace FleetDashClient.Models.Events;

public class OutgoingArmorEventArgs : EveLogEvent
{
    public OutgoingArmorEventArgs(string characterId, int amount, string pilot, string ship, string weapon) : base(
        characterId, amount, pilot, ship, weapon)
    {
    }
}