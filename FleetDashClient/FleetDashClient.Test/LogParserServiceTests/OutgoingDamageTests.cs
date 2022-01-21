using System;
using System.Collections.Generic;
using System.Text;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class OutgoingDamageTests
{
    private const string SingleDamageOutgoingLine =
        "[ 2022.01.15 23:15:57 ] (combat) <color=0xff00ffff><b>14</b> <color=0x77ffffff><font size=10>to</font> <b><color=0xffffffff>Anonymous Eve Player[TAGE](Exequror)</b><font size=10><color=0x77ffffff> - Acolyte II - Smashes";

    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

    public OutgoingDamageTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
    }
    
    [Fact]
    public void RaiseFileReadEvent_EmitCorrectTimestamp()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        var expected = new DateTimeOffset(2022, 01, 15, 23, 15, 57, TimeSpan.Zero);
        Assert.Equal(expected, emittedEvents[0].Timestamp);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingDamageNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitOutgoingDamageEvent()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Single(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal(14, emittedEvents[0].Amount);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("Exequror", emittedEvents[0].Ship);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("Acolyte II", emittedEvents[0].Weapon);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectApplication()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("Smashes", emittedEvents[0].Application);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<OutgoingDamageEvent>();

        _logParserService.OnOutgoingDamage += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].Pilot);
    }
}