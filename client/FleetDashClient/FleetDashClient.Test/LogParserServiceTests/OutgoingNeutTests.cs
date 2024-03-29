﻿using System;
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

public class OutgoingNeutTests
{
    private const string SingleNeutOutgoingLineWithOverview =
        "[ 2022.01.17 22:55:04 ] (combat) <color=0xff7fffff><b>44 GJ</b><color=0x77ffffff><font size=10> energy neutralized </font><b><color=0xffffffff><fontsize=12><color=0xFFFEBB64><b> <u>Vedmak</u></b></color></fontsize><fontsize=12><color=0xFFFEFF6F> [ALLIANCE]</color></fontsize> <fontsize=10><b>[CORP]</b></fontsize><fontsize=10> [Anonymous Eve Player]</fontsize><color=0xFFFFFFFF><b> -</b><color=0x77ffffff><font size=10> - Medium Energy Neutralizer II</font>";

    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

    private const string OverviewPath = @"TestData\ValidOverview.yaml";

    public OutgoingNeutTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
    }
    
    [Fact]
    public void RaiseFileReadEvent_EmitCorrectTimestamp()
    {
        var emittedEvents = new List<OutgoingNeutEvent>();

        _logParserService.OnOutgoingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutOutgoingLineWithOverview)));

        var expected = new DateTimeOffset(2022, 01, 17, 22, 55, 04, TimeSpan.Zero);
        Assert.Equal(expected, emittedEvents[0].Timestamp);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNeutNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<OutgoingNeutEvent>();

        _logParserService.OnOutgoingNeut += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutOutgoingLineWithOverview)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNeutWithOverview_ShouldEmitEvent()
    {
        var emittedEvents = new List<OutgoingNeutEvent>();

        _logParserService.OnOutgoingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutOutgoingLineWithOverview)));

        Assert.Single(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNeutWithOverview_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<OutgoingNeutEvent>();

        _logParserService.OnOutgoingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutOutgoingLineWithOverview)));

        Assert.Equal(44, emittedEvents[0].Amount);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNeutWithOverview_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<OutgoingNeutEvent>();

        _logParserService.OnOutgoingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutOutgoingLineWithOverview)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNeutWithOverview_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<OutgoingNeutEvent>();

        _logParserService.OnOutgoingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutOutgoingLineWithOverview)));

        Assert.Equal("Vedmak", emittedEvents[0].Ship);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNeutWithOverview_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<OutgoingNeutEvent>();

        _logParserService.OnOutgoingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutOutgoingLineWithOverview)));

        Assert.Equal("Medium Energy Neutralizer II", emittedEvents[0].Weapon);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNeutWithOverview_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<OutgoingNeutEvent>();

        _logParserService.OnOutgoingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", OverviewPath);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutOutgoingLineWithOverview)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].Pilot);
    }
}