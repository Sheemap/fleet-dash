using EveLogParser.Builder;
using EveLogParser.Models.Events;
using Microsoft.Extensions.Options;
using Serilog;

namespace EveLogParser;

internal class LogReaderService : ILogReaderService, IDisposable
{
    private readonly Dictionary<string, string?> _characterLogMap = new();
    private readonly Dictionary<string, int> _logProgress = new();
    private string _logDirectory;
    private CancellationTokenSource? _cancellationTokenSource;
    
    public LogReaderService(IOptionsMonitor<EveLogParserOptions> options)
    {
        _logDirectory = options.CurrentValue.LogDirectory;
        options.OnChange(ReloadConfig);
    }

    public event EventHandler<LogFileReadEventArgs>? OnFileRead;

    private void ReloadConfig(EveLogParserOptions options)
    {
        Log.Debug("Reloading log reader config");
        // Log directory needs to have changed to proceed
        if (options.LogDirectory == _logDirectory) return;

        _logDirectory = options.LogDirectory;

        // Only restart if currently running
        if (_cancellationTokenSource == null) return;
        
        Log.Information("Restarting log reader");
        Stop();
        Start();
    }

    public Task Start()
    {
        if (_cancellationTokenSource != null)
            throw new InvalidOperationException("Log reading already in progress. Cannot start the reader twice.");

        _cancellationTokenSource = new CancellationTokenSource();

        var token = _cancellationTokenSource.Token;

        return Task.Run(() =>
        {
            Log.Debug("Starting log watcher in directory {LogDirectory}", _logDirectory);
            using var watcher = new FileSystemWatcher(_logDirectory);

            watcher.NotifyFilter = NotifyFilters.LastWrite;

            watcher.Changed += OnChanged;

            watcher.Filter = "*.txt";
            watcher.EnableRaisingEvents = true;
            
            while (!token.IsCancellationRequested)
            {
                Thread.Sleep(1000);
            }
            
            Log.Information("Log reader shutting down.");
        });
    }

    public void Stop()
    {
        if (_cancellationTokenSource == null) throw new InvalidOperationException("Log reading is already stopped.");

        _cancellationTokenSource.Cancel();
        _cancellationTokenSource = null;
    }

    private string? GetCharacterId(string logFilePath)
    {
        Log.Debug("Getting character ID for path {FullPath}", logFilePath);
        string? charId;

        var cached = _characterLogMap.ContainsKey(logFilePath);
        if (!cached)
        {
            Log.Debug("Cache miss, parsing file");
            var logFileName = new FileInfo(logFilePath).Name;

            var fileNameParts = logFileName.Split('_');
            charId = fileNameParts.Length == 3 ? fileNameParts[2].Replace(".txt", "") : null;
            _characterLogMap.Add(logFilePath, charId);
            Log.Debug("Character ID {CharacterID} cached", charId);
        }
        else
        {
            charId = _characterLogMap.GetValueOrDefault(logFilePath);
            Log.Debug("Cache hit, character ID is {CharacterID}", charId);
        }

        return charId;
    }

    private void OnChanged(object sender, FileSystemEventArgs e)
    {
        var raiseEvent = OnFileRead;
        Log.Debug("Log read event fired.");
        if (raiseEvent == null) return;

        try
        {
            var charId = GetCharacterId(e.FullPath);

            if (charId == null) return;

            var position = _logProgress.GetValueOrDefault(e.FullPath, 0);

            Log.Debug("Opening {FullPath} for reading at position {FilePosition}", e.FullPath, position);
            using var fileStream = File.OpenRead(e.FullPath);

            fileStream.Position = position;

            var toRead = fileStream.Length - fileStream.Position;

            var readBytes = new byte[toRead];

            var readCount = fileStream.Read(readBytes, 0, Convert.ToInt32(toRead));

            _logProgress[e.FullPath] = position + readCount;

            Log.Debug("Read {Bytes} bytes, raising event.", readCount);
            var eventArgs = new LogFileReadEventArgs(charId, readBytes);
            raiseEvent(this, eventArgs);
        }
        catch (OverflowException)
        {
            _logProgress[e.FullPath] = 0;
        }
        catch (Exception ex)
        {
            Log.Error(ex, "Exception in log reader occured!");
        }
    }

    public void Dispose()
    {
        _cancellationTokenSource?.Dispose();
    }
}