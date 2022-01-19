using System.IO;
using FleetDashClient.Services;
using Moq;
using Xunit;
using YamlDotNet.Core;

namespace FleetDashClient.Test.LogParserServiceTests;

public class MiscTests
{
    private readonly ILogParserService _logParserService;
    private readonly Mock<ILogReaderService> _logReaderMock;

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
        Assert.ThrowsAny<YamlException>(
            () => _logParserService.StartWatchingCharacter("123", overviewYaml));
    }
}