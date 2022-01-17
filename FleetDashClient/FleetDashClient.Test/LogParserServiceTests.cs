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

        private readonly string _singleDamageIncomingLine = "[ 2022.01.15 23:16:37 ] (combat) <color=0xffcc0000><b>164</b> <color=0x77ffffff><font size=10>from</font> <b><color=0xffffffff>Anonymous Eve Playeer[TAGE](Harbinger)</b><font size=10><color=0x77ffffff> - Focused Medium Beam Laser II - Hits";

        public LogParserServiceTests()
        {
            _logReaderMock = new Mock<ILogReaderService>();

            _logParserService = new LogParserService(_logReaderMock.Object);
        }

        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitReceivedDamageEvent()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseDamageReceivedEvent += (_, e) => emittedEvents.Add(e);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(_singleDamageIncomingLine)));

            Assert.Single(emittedEvents);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectCharacterID()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseDamageReceivedEvent += (_, e) => emittedEvents.Add(e);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(_singleDamageIncomingLine)));

            Assert.Equal("123", emittedEvents[0].CharacterID);
        }
        
        [Fact]
        public void RaiseFileReadEvent_IncomingDamage_ShouldEmitCorrectAmount()
        {
            var emittedEvents = new List<IncomingDamageEventArgs>();

            _logParserService.RaiseDamageReceivedEvent += (_, e) => emittedEvents.Add(e);

            // Emit the event
            _logReaderMock.Raise(x => x.RaiseFileReadEvent += null,
                new LogFileReadEventArgs("123", Encoding.UTF8.GetBytes(_singleDamageIncomingLine)));

            Assert.Equal(164, emittedEvents[0].Amount);
        }
    }
}
