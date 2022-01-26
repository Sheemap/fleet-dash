using System.Runtime.InteropServices;
using ElectronNET.API;
using ElectronNET.API.Entities;
using FleetDashClient.Configuration;
using FleetDashClient.Data;
using FleetDashClient.Models.Events;
using FleetDashClient.Services;
using Microsoft.Extensions.Options;

namespace FleetDashClient.ViewModels;

public class LogViewModel

{
    public event EventHandler<PathUpdatedEventArgs> OnLogDirectoryUpdated;
    public event EventHandler<PathUpdatedEventArgs> OnOverviewFileUpdated;
    
    public string LogDirectory { get; set; }
    public string OverviewFile { get; set; }
    
    private readonly IOptionsMonitor<Models.Configuration> _configuration;

    private readonly JsonConfigurationManager _configManager;

    public LogViewModel(IOptionsMonitor<Models.Configuration> config, JsonConfigurationManager configManager)
    {
        _configuration = config;
        _configManager = configManager;

        LogDirectory = _configuration.CurrentValue.LogDirectory;
        OverviewFile = _configuration.CurrentValue.OverviewPath;
        config.OnChange(UpdateConfig);
    }
    
    private void UpdateConfig(Models.Configuration config)
    {
        // TODO: Why does updating these props not update the UI?
        // blazor wack, learn it better
        LogDirectory = config.LogDirectory;
        OverviewFile = config.OverviewPath;
    }


    public async Task UpdateLogDirectory()
    {
        var mainWindow = Electron.WindowManager.BrowserWindows.First();
        var options = new OpenDialogOptions {
            Properties = new[] {
                OpenDialogProperty.openDirectory
            }
        };

        string[] folder = await Electron.Dialog.ShowOpenDialogAsync(mainWindow, options);
        if (folder.Length > 0)
        {
            var config = new Models.Configuration
            {
                LogDirectory = folder[0]
            };
            _configManager.UpdateConfiguration(config);
        }
    }

    public async Task UpdateOverview()
    {
        var mainWindow = Electron.WindowManager.BrowserWindows.First();
        var options = new OpenDialogOptions {
            Properties = new[] {
                OpenDialogProperty.openFile
            }
        };

        string[] file = await Electron.Dialog.ShowOpenDialogAsync(mainWindow, options);
        if (file.Length > 0)
        {
            var config = new Models.Configuration
            {
                OverviewPath = file[0]
            };
            _configManager.UpdateConfiguration(config);
        }
    }
}