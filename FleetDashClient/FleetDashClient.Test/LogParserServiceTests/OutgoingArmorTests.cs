using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class OutgoingArmorTests
{
    private const string SingleArmorOutgoingLineWithOverview =
        "[ 2022.01.17 20:28:44 ] (combat) <color=0xffccff66><b>416</b><color=0x77ffffff><font size=10> remote armor repaired to </font><b><color=0xffffffff><fontsize=12><color=0xFFFEBB64><b> <u>Hurricane</u></b></color></fontsize><fontsize=12><color=0xFFFEFF6F> [ALLIANCE_TAG]</color></fontsize> <fontsize=10><b>[CORP_TAG]</b></fontsize><fontsize=10> [Anonymous Eve Player]</fontsize><color=0xFFFFFFFF><b> -</b><color=0x77ffffff><font size=10> - Medium Remote Armor Repairer II</font>";

    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

    private readonly string _overviewSettings;

    public OutgoingArmorTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);

        using var streamReader = new StreamReader(@"TestData\ValidOverview.yaml");
        _overviewSettings = streamReader.ReadToEnd();
    }
    
    [Fact]
    public void RaiseFileReadEvent_EmitCorrectTimestamp()
    {
        var emittedEvents = new List<OutgoingArmorEvent>();

        _logParserService.OnOutgoingArmor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorOutgoingLineWithOverview)));

        var expected = new DateTimeOffset(2022, 01, 17, 20, 28, 44, TimeSpan.Zero);
        Assert.Equal(expected, emittedEvents[0].Timestamp);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingArmorNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<OutgoingArmorEvent>();

        _logParserService.OnOutgoingArmor += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorOutgoingLineWithOverview)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingArmorWithOverview_ShouldEmitEvent()
    {
        var emittedEvents = new List<OutgoingArmorEvent>();

        _logParserService.OnOutgoingArmor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorOutgoingLineWithOverview)));

        Assert.Single(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingArmorWithOverview_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<OutgoingArmorEvent>();

        _logParserService.OnOutgoingArmor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorOutgoingLineWithOverview)));

        Assert.Equal(416, emittedEvents[0].Amount);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingArmorWithOverview_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<OutgoingArmorEvent>();

        _logParserService.OnOutgoingArmor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorOutgoingLineWithOverview)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingArmorWithOverview_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<OutgoingArmorEvent>();

        _logParserService.OnOutgoingArmor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorOutgoingLineWithOverview)));

        Assert.Equal("Hurricane", emittedEvents[0].Ship);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingArmorWithOverview_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<OutgoingArmorEvent>();

        _logParserService.OnOutgoingArmor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorOutgoingLineWithOverview)));

        Assert.Equal("Medium Remote Armor Repairer II", emittedEvents[0].Weapon);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingArmorWithOverview_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<OutgoingArmorEvent>();

        _logParserService.OnOutgoingArmor += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorOutgoingLineWithOverview)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].Pilot);
    }
}