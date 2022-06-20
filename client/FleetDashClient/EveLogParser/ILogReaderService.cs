using EveLogParser.Models.Events;

namespace EveLogParser;

public interface ILogReaderService
{
    event EventHandler<LogFileReadEventArgs> OnFileRead;

    Task Start();
    void Stop();
}