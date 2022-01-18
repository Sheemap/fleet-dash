using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace FleetDashClient.Test
{
    public class LogParserServiceTests
    {
        private readonly Mock<ILogReaderService> _logReaderMock;
        private readonly ILogParserService _logParserService;

        private const string SingleDamageIncomingLine =
            "[ 2022.01.15 23:16:37 ] (combat) <color=0xffcc0000><b>164</b> <color=0x77ffffff><font size=10>from</font> <b><color=0xffffffff>Anonymous Eve Player[TAGE](Harbinger)</b><font size=10><color=0x77ffffff> - Focused Medium Beam Laser II - Penetrates";
        private const string SingleDamageOutgoingLine =
            "[ 2022.01.15 23:15:57 ] (combat) <color=0xff00ffff><b>14</b> <color=0x77ffffff><font size=10>to</font> <b><color=0xffffffff>Anonymous Eve Player[TAGE](Exequror)</b><font size=10><color=0x77ffffff> - Acolyte II - Smashes";

        public LogParserServiceTests()
        {
            _logReaderMock = new Mock<ILogReaderService>();

            _logParserService = new LogParserService(_logReaderMock.Object);
        }
        
        [Fact]
        public void RaiseFileReadEvent_NonRegisteredCharacterIncomingDamage_ShouldNotEmitEvent()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Empty(emittedEvents);
        }

        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitIncomingDamageEvent()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
            _logParserService.StartWatchingCharacter("123", null);
            _logParserService.StartWatchingCharacter("123", null);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Single(emittedEvents);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectCharacterID()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
            _logParserService.StartWatchingCharacter("123", null);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Equal("123", emittedEvents[0].CharacterId);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectAmount()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
            _logParserService.StartWatchingCharacter("123", null);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Equal(164, emittedEvents[0].Amount);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectName()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
            _logParserService.StartWatchingCharacter("123", null);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Equal("Anonymous Eve Player", emittedEvents[0].FromName);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectTag()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
            _logParserService.StartWatchingCharacter("123", null);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Equal("TAGE", emittedEvents[0].FromTag);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectShip()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
            _logParserService.StartWatchingCharacter("123", null);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Equal("Harbinger", emittedEvents[0].FromShip);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectWeapon()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
            _logParserService.StartWatchingCharacter("123", null);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Equal("Focused Medium Beam Laser II", emittedEvents[0].Weapon);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectApplication()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseIncomingDamageEvent += (_, e) => emittedEvents.Add(e);
            _logParserService.StartWatchingCharacter("123", null);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(SingleDamageIncomingLine)));

            Assert.Equal("Penetrates", emittedEvents[0].Application);
        }
        
        [Fact]
        public void RaiseFileReadEvent_NonRegisteredCharacterOutgoingDamage_ShouldNotEmitEvent()
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
}
