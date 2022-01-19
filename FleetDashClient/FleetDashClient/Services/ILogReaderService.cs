using FleetDashClient.Models.Events;

namespace FleetDashClient.Services;

public interface ILogReaderService
{
    event EventHandler<LogFileReadEventArgs> RaiseFileReadEvent;

    Task Start();
    void Stop();
}