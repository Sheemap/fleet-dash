﻿using System.Collections.Generic;
using System.IO;
using System.Text;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class IncomingArmorTests
{
    private readonly Mock<ILogReaderService> _logReaderMock;
    private readonly ILogParserService _logParserService;
    
    private readonly string _overviewSettings;
    
    private const string SingleArmorIncomingLineWithOverview = 
        "[ 2022.01.17 22:57:14 ] (combat) <color=0xffccff66><b>416</b><color=0x77ffffff><font size=10> remote armor repaired by </font><b><color=0xffffffff><fontsize=12><color=0xFFFEBB64><b> <u>Augoror</u></b></color></fontsize><fontsize=12><color=0xFFFEFF6F> [ALLIANCE]</color></fontsize> <fontsize=10><b>[CORP]</b></fontsize><fontsize=10> [Anonymous Eve Player]</fontsize><color=0xFFFFFFFF><b> -</b><color=0x77ffffff><font size=10> - Medium Remote Armor Repairer II</font>";
    
    public IncomingArmorTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
        
        using var streamReader = new StreamReader(@"TestData\ValidOverview.yaml");
        _overviewSettings = streamReader.ReadToEnd();
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingArmorNonRegisteredCharacter_ShouldNotEmitEvent()
    {
        var emittedEvents = new List<IncomingArmorEventArgs>();

        _logParserService.RaiseIncomingArmorEvent += (_, e) => emittedEvents.Add(e);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorIncomingLineWithOverview)));

        Assert.Empty(emittedEvents);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingArmorWithOverview_ShouldEmitEvent()
    {
        var emittedEvents = new List<IncomingArmorEventArgs>();

        _logParserService.RaiseIncomingArmorEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorIncomingLineWithOverview)));

        Assert.Single(emittedEvents);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingArmorWithOverview_ShouldEmitCorrectAmount()
    {
        var emittedEvents = new List<IncomingArmorEventArgs>();

        _logParserService.RaiseIncomingArmorEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorIncomingLineWithOverview)));

        Assert.Equal(416, emittedEvents[0].Amount);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingArmorWithOverview_ShouldEmitCorrectCharacterId()
    {
        var emittedEvents = new List<IncomingArmorEventArgs>();

        _logParserService.RaiseIncomingArmorEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorIncomingLineWithOverview)));

        Assert.Equal("123", emittedEvents[0].CharacterId);
    }
    [Fact]
    public void RaiseFileReadEvent_IncomingArmorWithOverview_ShouldEmitCorrectShip()
    {
        var emittedEvents = new List<IncomingArmorEventArgs>();

        _logParserService.RaiseIncomingArmorEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorIncomingLineWithOverview)));

        Assert.Equal("Augoror", emittedEvents[0].FromShip);
    }
    
    [Fact]
    public void RaiseFileReadEvent_IncomingArmorWithOverview_ShouldEmitCorrectWeapon()
    {
        var emittedEvents = new List<IncomingArmorEventArgs>();

        _logParserService.RaiseIncomingArmorEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorIncomingLineWithOverview)));

        Assert.Equal("Medium Remote Armor Repairer II", emittedEvents[0].Weapon);
    }
    [Fact]
    public void RaiseFileReadEvent_IncomingArmorWithOverview_ShouldEmitCorrectName()
    {
        var emittedEvents = new List<IncomingArmorEventArgs>();

        _logParserService.RaiseIncomingArmorEvent += (_, e) => emittedEvents.Add(e);
        _logParserService.StartWatchingCharacter("123", _overviewSettings);

        // Emit the event
        _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
            new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleArmorIncomingLineWithOverview)));

        Assert.Equal("Anonymous Eve Player", emittedEvents[0].FromName);
    }
}