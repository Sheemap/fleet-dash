using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using EveLogParser;
using EveLogParser.Models.Events;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class IncomingCapacitorTests
{
    private const string SingleCapacitorIncomingLineWithOverview =
        "[ 2022.01.17 22:57:13 ] (combat) <color=0xffccff66><b>351</b><color=0x77ffffff><font size=10> remote capacitor transmitted by </font><b><color=0xffffffff><fontsize=12><color=0xFFFEBB64><b> <u>Augoror</u></b></color></fontsize><fontsize=12><color=0xFFFEFF6F> [ALLIANCE]</color></fontsize> <fontsize=10><b>[CORP]</b></fontsize><fontsize=10> [Anonymous Eve Player]</fontsize><color=0xFFFFFFFF><b> -</b><color=0x77ffffff><font size=10> - Medium Remote Capacitor Transmitter II</font>";

    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

    private const string OverviewPath = @"TestData\ValidOverview.yaml";

    public IncomingCapacitorTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
    }
    
    [Fact]
    public void RaiseFileReadEvent_EmitCorrectTimestamp()
    {
        var emittedEvents = new List<IncomingCapacitorEvent>();

        _logParserService.OnIncomingCapacitor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleCapacitorIncomingLineWithOverview)));

        var expected = new DateTimeOffset(2022, 01, 17, 22, 57, 13, TimeSpan.Zero);
        Assert.Equal(expected, emittedEvents[0].Timestamp);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingCapacitorNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<IncomingCapacitorEvent>();

        _logParserService.OnIncomingCapacitor += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleCapacitorIncomingLineWithOverview)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingCapacitorWithOverview_ShouldEmitEvent()
    {
        var emittedEvents = new List<IncomingCapacitorEvent>();

        _logParserService.OnIncomingCapacitor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleCapacitorIncomingLineWithOverview)));

        Assert.Single(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingCapacitorWithOverview_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<IncomingCapacitorEvent>();

        _logParserService.OnIncomingCapacitor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleCapacitorIncomingLineWithOverview)));

        Assert.Equal(351, emittedEvents[0].Amount);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingCapacitorWithOverview_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<IncomingCapacitorEvent>();

        _logParserService.OnIncomingCapacitor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleCapacitorIncomingLineWithOverview)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingCapacitorWithOverview_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<IncomingCapacitorEvent>();

        _logParserService.OnIncomingCapacitor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleCapacitorIncomingLineWithOverview)));

        Assert.Equal("Augoror", emittedEvents[0].Ship);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingCapacitorWithOverview_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<IncomingCapacitorEvent>();

        _logParserService.OnIncomingCapacitor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleCapacitorIncomingLineWithOverview)));

        Assert.Equal("Medium Remote Capacitor Transmitter II", emittedEvents[0].Weapon);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingCapacitorWithOverview_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<IncomingCapacitorEvent>();

        _logParserService.OnIncomingCapacitor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleCapacitorIncomingLineWithOverview)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].Pilot);
    }
}