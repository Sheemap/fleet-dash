using System;
using System.Collections.Generic;
using System.Text;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class IncomingDamageTests
{
    private const string SingleDamageIncomingLine =
        "[ 2022.01.15 23:16:37 ] (combat) <color=0xffcc0000><b>164</b> <color=0x77ffffff><font size=10>from</font> <b><color=0xffffffff>Anonymous Eve Player[TAGE](Harbinger)</b><font size=10><color=0x77ffffff> - Focused Medium Beam Laser II - Penetrates";

    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

    public IncomingDamageTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
    }
    
    [Fact]
    public void RaiseFileReadEvent_EmitCorrectTimestamp()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        var expected = new DateTimeOffset(2022, 01, 15, 23, 16, 37, TimeSpan.Zero);
        Assert.Equal(expected, emittedEvents[0].Timestamp);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamageNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitIncomingDamageEvent()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Single(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectCharacterID()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal(164, emittedEvents[0].Amount);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].Pilot);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("Harbinger", emittedEvents[0].Ship);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("Focused Medium Beam Laser II", emittedEvents[0].Weapon);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectApplication()
    {
        var emittedEvents = new List<IncomingDamageEvent>();

        _logParserService.OnIncomingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

        Assert.Equal("Penetrates", emittedEvents[0].Application);
    }
}