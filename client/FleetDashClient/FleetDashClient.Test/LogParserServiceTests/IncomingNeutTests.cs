using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class IncomingNeutTests
{
    private const string SingleNeutIncomingLineWithOverview =
        "[ 2022.01.17 22:55:04 ] (combat) <color=0xffe57f7f><b>44 GJ</b><color=0x77ffffff><font size=10> energy neutralized </font><b><color=0xffffffff><fontsize=12><color=0xFFFEBB64><b> <u>Vedmak</u></b></color></fontsize><fontsize=12><color=0xFFFEFF6F> [ALLIANCE]</color></fontsize> <fontsize=10><b>[CORP]</b></fontsize><fontsize=10> [Anonymous Eve Player]</fontsize><color=0xFFFFFFFF><b> -</b><color=0x77ffffff><font size=10> - Medium Energy Neutralizer II</font>";

    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

    private readonly string _overviewSettings;

    public IncomingNeutTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);

        using var streamReader = new StreamReader(@"TestData\ValidOverview.yaml");
        _overviewSettings = streamReader.ReadToEnd();
    }
    
    [Fact]
    public void RaiseFileReadEvent_EmitCorrectTimestamp()
    {
        var emittedEvents = new List<IncomingNeutEvent>();

        _logParserService.OnIncomingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutIncomingLineWithOverview)));

        var expected = new DateTimeOffset(2022, 01, 17, 22, 55, 04, TimeSpan.Zero);
        Assert.Equal(expected, emittedEvents[0].Timestamp);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNeutNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<IncomingNeutEvent>();

        _logParserService.OnIncomingNeut += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutIncomingLineWithOverview)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNeutWithOverview_ShouldEmitEvent()
    {
        var emittedEvents = new List<IncomingNeutEvent>();

        _logParserService.OnIncomingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutIncomingLineWithOverview)));

        Assert.Single(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNeutWithOverview_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<IncomingNeutEvent>();

        _logParserService.OnIncomingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutIncomingLineWithOverview)));

        Assert.Equal(44, emittedEvents[0].Amount);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNeutWithOverview_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<IncomingNeutEvent>();

        _logParserService.OnIncomingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutIncomingLineWithOverview)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNeutWithOverview_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<IncomingNeutEvent>();

        _logParserService.OnIncomingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutIncomingLineWithOverview)));

        Assert.Equal("Vedmak", emittedEvents[0].Ship);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNeutWithOverview_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<IncomingNeutEvent>();

        _logParserService.OnIncomingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutIncomingLineWithOverview)));

        Assert.Equal("Medium Energy Neutralizer II", emittedEvents[0].Weapon);
    }

    [Fact]
    public void RaiseFileReadEvent_IncomingNeutWithOverview_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<IncomingNeutEvent>();

        _logParserService.OnIncomingNeut += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNeutIncomingLineWithOverview)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].Pilot);
    }
}