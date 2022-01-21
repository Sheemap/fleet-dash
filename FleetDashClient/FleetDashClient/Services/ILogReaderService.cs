using FleetDashClient.Models.Events;

namespace FleetDashClient.Services;

public interface ILogReaderService
{
    event EventHandler<LogFileReadEventArgs> OnFileRead;

    Task Start();
    void Stop();
}