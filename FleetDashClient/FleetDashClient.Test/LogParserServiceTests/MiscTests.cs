using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Moq;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace FleetDashClient.Test.LogParserServiceTests;

public class MiscTests
{
    private readonly Mock<ILogReaderService> _logReaderMock;
    private readonly ILogParserService _logParserService;

    public MiscTests()
    {
        _logReaderMock = new Mock<ILogReaderService>();

        _logParserService = new LogParserService(_logReaderMock.Object);
    }

    [Fact]
    public void StartWatchingCharacter_ValidYaml_ShouldNotThrow()
    {
        // Arrange
        using var streamReader = new StreamReader(@"TestData\ValidOverview.yaml");
        var overviewYaml = streamReader.ReadToEnd();

        // Act
        _logParserService.StartWatchingCharacter("123", overviewYaml);
    }
    
    [Fact]
    public void StartWatchingCharacter_InvalidYaml_ShouldThrow()
    {
        // Arrange
        using var streamReader = new StreamReader(@"TestData\InvalidOverview.yaml");
        var overviewYaml = streamReader.ReadToEnd();

        // Act
        Assert.ThrowsAny<YamlDotNet.Core.YamlException>(
            () => _logParserService.StartWatchingCharacter("123", overviewYaml));
        
    }
}

