using ElectronNET.API;
using ElectronNET.API.Entities;
using EventAggregator.Blazor;
using FleetDashClient.Configuration;
using FleetDashClient.Models.Events;
using Microsoft.Extensions.Options;
using Serilog;

namespace FleetDashClient.ViewModels;

public class LogViewModel

{
    public event EventHandler<PathUpdatedEventArgs> OnLogDirectoryUpdated;
    
    public string LogDirectory { get; set; }
    public string OverviewPath { get; set; }
    
    private readonly IOptionsMonitor<Models.Configuration> _configuration;
    private readonly JsonConfigurationManager _configManager;
    private readonly IEventAggregator _eventAggregator;

    public LogViewModel(
        IOptionsMonitor<Models.Configuration> config, 
        JsonConfigurationManager configManager, 
        IEventAggregator eventAggregator)
    {
        _configuration = config;
        _configManager = configManager;
        _eventAggregator = eventAggregator;

        LogDirectory = _configuration.CurrentValue.LogDirectory;
        config.OnChange(UpdateConfig);
    }
    
    private void UpdateConfig(Models.Configuration config)
    {
        LogDirectory = config.LogDirectory;
        _eventAggregator.PublishAsync(new PathUpdatedEventArgs(string.Empty));
    }


    public async Task UpdateLogDirectory()
    {
        var mainWindow = Electron.WindowManager.BrowserWindows.First();
        var options = new OpenDialogOptions {
            Properties = new[] {
                OpenDialogProperty.openDirectory
            },
            Title = "Select Log Directory",
        };

        Log.Debug("Prompting for folder selection");
        string[] folder = await Electron.Dialog.ShowOpenDialogAsync(mainWindow, options);
        Log.Debug("Log directory selected: {@Folder}", folder);
        if (folder.Length > 0)
        {
            var config = new Models.Configuration
            {
                LogDirectory = folder[0]
            };
            _configManager.UpdateConfiguration(config);
        }
    }
}