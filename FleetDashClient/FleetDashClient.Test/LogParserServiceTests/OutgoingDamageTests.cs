using System.Collections.Generic;
using System.Text;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class OutgoingDamageTests
{
    private readonly Mock<ILogReaderService> _logReaderMock;
    private readonly ILogParserService _logParserService;
    
    private const string SingleDamageOutgoingLine =
        "[ 2022.01.15 23:15:57 ] (combat) <color=0xff00ffff><b>14</b> <color=0x77ffffff><font size=10>to</font> <b><color=0xffffffff>Anonymous Eve Player[TAGE](Exequror)</b><font size=10><color=0x77ffffff> - Acolyte II - Smashes";
    
    public OutgoingDamageTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamageNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Empty(emittedEvents);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitOutgoingDamageEvent()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Single(emittedEvents);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal(14, emittedEvents[0].Amount);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectTag()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("TAGE", emittedEvents[0].ToTag);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("Exequror", emittedEvents[0].ToShip);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("Acolyte II", emittedEvents[0].Weapon);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectApplication()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("Smashes", emittedEvents[0].Application);
    }
    
    [Fact]
    public void RaiseFileReadEvent_OutgoingDamage_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<OutgoingDamageEventArgs>();

        _logParserService.RaiseOutgoingDamageEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", null);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageOutgoingLine)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].ToName);
    }
}