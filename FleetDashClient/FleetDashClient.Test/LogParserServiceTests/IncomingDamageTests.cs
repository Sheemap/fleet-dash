﻿using System.Collections.Generic;
using System.Text;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class IncomingDamageTests
{
    private readonly Mock<ILogReaderService> _logReaderMock;
    private readonly ILogParserService _logParserService;
    
    private const string SingleDamageIncomingLine =
        "[ 2022.01.15 23:16:37 ] (combat) <color=0xffcc0000><b>164</b> <color=0x77ffffff><font size=10>from</font> <b><color=0xffffffff>Anonymous Eve Player[TAGE](Harbinger)</b><font size=10><color=0x77ffffff> - Focused Medium Beam Laser II - Penetrates";
    
    public IncomingDamageTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingDamageNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitIncomingDamageEvent()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Single(emittedEvents);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectCharacterID()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal(164, emittedEvents[0].Amount);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].FromName);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectTag()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("TAGE", emittedEvents[0].FromTag);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("Harbinger", emittedEvents[0].FromShip);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("Focused Medium Beam Laser II", emittedEvents[0].Weapon);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectApplication()
    {
        var emittedEvents = new List<IncomingDamageEventArgs>();

        _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("Penetrates", emittedEvents[0].Application);
    }
}