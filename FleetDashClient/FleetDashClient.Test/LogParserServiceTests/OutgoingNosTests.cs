using System.Collections.Generic;
using System.IO;
using System.Text;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class OutgoingNosTests
{
    private const string SingleNosOutgoingLineWithOverview =
        "[ 2022.01.15 22:14:11 ] (combat) <color=0xffe57f7f><b>+20 GJ</b><color=0x77ffffff><font size=10> energy drained from </font><b><color=0xffffffff><fontsize=12><color=0xFFFEBB64><b> <u>Bifrost</u></b></color></fontsize><fontsize=12><color=0xFFFEFF6F> [ALLIANCE]</color></fontsize> <fontsize=10><b>[CORP]</b></fontsize><fontsize=10> [Anonymous Eve Player]</fontsize><color=0xFFFFFFFF><b> -</b><color=0x77ffffff><font size=10> - Small Ghoul Compact Energy Nosferatu</font>";

    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

    private readonly string _overviewSettings;

    public OutgoingNosTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);

        using var streamReader = new StreamReader(@"TestData\ValidOverview.yaml");
        _overviewSettings = streamReader.ReadToEnd();
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNosNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<OutgoingNosEvent>();

        _logParserService.OnOutgoingNos += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosOutgoingLineWithOverview)));

        Assert.Empty(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNosWithOverview_ShouldEmitEvent()
    {
        var emittedEvents = new List<OutgoingNosEvent>();

        _logParserService.OnOutgoingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosOutgoingLineWithOverview)));

        Assert.Single(emittedEvents);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNosWithOverview_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<OutgoingNosEvent>();

        _logParserService.OnOutgoingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosOutgoingLineWithOverview)));

        Assert.Equal(20, emittedEvents[0].Amount);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNosWithOverview_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<OutgoingNosEvent>();

        _logParserService.OnOutgoingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosOutgoingLineWithOverview)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNosWithOverview_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<OutgoingNosEvent>();

        _logParserService.OnOutgoingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosOutgoingLineWithOverview)));

        Assert.Equal("Bifrost", emittedEvents[0].Ship);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNosWithOverview_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<OutgoingNosEvent>();

        _logParserService.OnOutgoingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosOutgoingLineWithOverview)));

        Assert.Equal("Small Ghoul Compact Energy Nosferatu", emittedEvents[0].Weapon);
    }

    [Fact]
    public void RaiseFileReadEvent_OutgoingNosWithOverview_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<OutgoingNosEvent>();

        _logParserService.OnOutgoingNos += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.OnFileRead += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleNosOutgoingLineWithOverview)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].Pilot);
    }
}