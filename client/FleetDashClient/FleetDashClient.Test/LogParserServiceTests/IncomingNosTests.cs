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

public class IncomingNosTests
{
    private const string SingleNosIncomingLineWithOverview =
        "[ 2022.01.15 22:14:11 ] (combat) <color=0xffe57f7f><b>-20 GJ</b><color=0x77ffffff><font size=10> energy drained to </font><b><color=0xffffffff><fontsize=12><color=0xFFFEBB64><b> <u>Bifrost</u></b></color></fontsize><fontsize=12><color=0xFFFEFF6F> [ALLIANCE]</color></fontsize> <fontsize=10><b>[CORP]</b></fontsize><fontsize=10> [Anonymous Eve Player]</fontsize><color=0xFFFFFFFF><b> -</b><color=0x77ffffff><font size=10> - Small Ghoul Compact Energy Nosferatu</font>";

    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

    private const string OverviewPath = @"TestData\ValidOverview.yaml";

    public IncomingNosTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
    }
    
    [Fact]
    public void RaiseFileReadEvent_EmitCorrectTimestamp()
    {
        var emittedEvents = new List<IncomingNosEvent>();

        _logParserService.OnIncomingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosIncomingLineWithOverview)));

        var expected = new DateTimeOffset(2022, 01, 15, 22, 14, 11, TimeSpan.Zero);
        Assert.Equal(expected, emittedEvents[0].Timestamp);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNosNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<IncomingNosEvent>();

        _logParserService.OnIncomingNos += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosIncomingLineWithOverview)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNosWithOverview_ShouldEmitEvent()
    {
        var emittedEvents = new List<IncomingNosEvent>();

        _logParserService.OnIncomingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosIncomingLineWithOverview)));

        Assert.Single(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNosWithOverview_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<IncomingNosEvent>();

        _logParserService.OnIncomingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosIncomingLineWithOverview)));

        Assert.Equal(20, emittedEvents[0].Amount);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNosWithOverview_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<IncomingNosEvent>();

        _logParserService.OnIncomingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosIncomingLineWithOverview)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNosWithOverview_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<IncomingNosEvent>();

        _logParserService.OnIncomingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosIncomingLineWithOverview)));

        Assert.Equal("Bifrost", emittedEvents[0].Ship);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNosWithOverview_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<IncomingNosEvent>();

        _logParserService.OnIncomingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosIncomingLineWithOverview)));

        Assert.Equal("Small Ghoul Compact Energy Nosferatu", emittedEvents[0].Weapon);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNosWithOverview_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<IncomingNosEvent>();

        _logParserService.OnIncomingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosIncomingLineWithOverview)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].Pilot);
    }
}