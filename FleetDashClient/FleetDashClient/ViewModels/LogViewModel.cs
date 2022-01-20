using System.Runtime.InteropServices;
using ElectronNET.API;
using ElectronNET.API.Entities;
using FleetDashClient.Data;
using FleetDashClient.Services;

namespace FleetDashClient.ViewModels;

public class LogViewModel
{
    public string LogDirectory { get; set; }
    public string? Overview { get; set; }

    private readonly DataContext _dbContext;

    public LogViewModel(DataContext dbContext)
    {
        _dbContext = dbContext;
        InitializeLogSettings();
        
    }

    private void InitializeLogSettings()
    {
        var config = _dbContext.Configurations.FirstOrDefault();

        if (config != null && !string.IsNullOrWhiteSpace(config.LogDirectory))
        {
            LogDirectory = config.LogDirectory;
        }

        if (config != null && !string.IsNullOrWhiteSpace(config.Overview))
        {
            Overview = Path.GetFileName(config.Overview);
        }
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
            LogDirectory = folder[0];
            var config = _dbContext.Configurations.First();
            config.LogDirectory = LogDirectory;
            await _dbContext.SaveChangesAsync();
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
            Overview = Path.GetFileName(file[0]);
            var config = _dbContext.Configurations.First();
            config.Overview = file[0];
            await _dbContext.SaveChangesAsync();
        }
    }
}