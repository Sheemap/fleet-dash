using FleetDashClient.Models.Events;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Services
{
    public class LogReaderService : ILogReaderService
    {
        private string _logDirectory;
        private CancellationTokenSource cancellationTokenSource;
        private readonly Dictionary<string, string> _characterLogMap = new Dictionary<string, string>();
        private readonly Dictionary<string, int> _logProgress = new Dictionary<string, int>();

        public event EventHandler<LogFileReadEventArgs> RaiseFileReadEvent;


        public LogReaderService(string logDirectory)
        {
            _logDirectory = logDirectory;
        }

        public Task Start()
        {
            if (cancellationTokenSource != null)
            {
                throw new InvalidOperationException("Log reading already in progress. Cannot start the reader twice.");
            }

            cancellationTokenSource = new CancellationTokenSource();

            var token = cancellationTokenSource.Token;

            return Task.Run(() =>
            {
                using var watcher = new FileSystemWatcher(_logDirectory);

                watcher.NotifyFilter = NotifyFilters.LastWrite;

                watcher.Changed += OnChanged;

                watcher.Filter = "*.txt";
                watcher.EnableRaisingEvents = true;

                while (!token.IsCancellationRequested)
                {
                    Task.Delay(1);
                }
            });
        }

        public void Stop()
        {
            if (cancellationTokenSource == null)
            {
                throw new InvalidOperationException("Log reading is already stopped.");
            }

            cancellationTokenSource.Cancel();
        }

        private string GetCharacterID(string logFilePath)
        {
            string charId;

            var cached = _characterLogMap.ContainsKey(logFilePath);
            if (!cached)
            {
                var logFileName = new FileInfo(logFilePath).Name;

                var fileNameParts = logFileName.Split('_');
                if (fileNameParts.Length == 3)
                {
                    charId = fileNameParts[2].Replace(".txt", "");
                }
                else
                {
                    charId = null;
                }
                _characterLogMap.Add(logFilePath, charId);
            }
            else
            {
                charId = _characterLogMap.GetValueOrDefault(logFilePath);
            }

            return charId;
        }

        private void OnChanged(object sender, FileSystemEventArgs e)
        {
            var raiseEvent = RaiseFileReadEvent;
            if (raiseEvent != null)
            {
                var charId = GetCharacterID(e.FullPath);

                if (charId == null)
                {
                    return;
                }

                var position = _logProgress.GetValueOrDefault(e.FullPath, 0);

                using var fileStream = File.OpenRead(e.FullPath);

                fileStream.Position = position;

                var toRead = fileStream.Length - fileStream.Position;

                var readBytes = new byte[toRead];

                var readCount = fileStream.Read(readBytes, 0, Convert.ToInt32(toRead));

                _logProgress[e.FullPath] = position + readCount;

                var eventArgs = new LogFileReadEventArgs(charId, readBytes);
                raiseEvent(this, eventArgs);
            }
        }
    }
}
