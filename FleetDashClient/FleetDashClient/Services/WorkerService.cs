using ElectronNET.API;
using FleetDashClient.Data;

namespace FleetDashClient.Services;

public class WorkerService : BackgroundService
{
    private readonly DataContext _dbContext;
    
    private bool _handlersInitialized;
    
    public WorkerService(DataContext dbContext)
    {
        _dbContext = dbContext;
    }
    
    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
        while (!stoppingToken.IsCancellationRequested)
        {
            AddEventHandlers();
            await Task.Delay(1000, stoppingToken);
        }
    }

    private void AddEventHandlers()
    {
        if (_handlersInitialized) return;
        
        var mainWindow = Electron.WindowManager.BrowserWindows.FirstOrDefault();
        if (mainWindow == null) return;
        
        mainWindow.OnMove += UpdateWindowPosition;
        mainWindow.OnResize += UpdateWindowSize;
        
        _handlersInitialized = true;
    }

    private async void UpdateWindowPosition()
    {
        var mainWindow = Electron.WindowManager.BrowserWindows.First();
        var windowPosition = await mainWindow.GetPositionAsync();

        var config = _dbContext.Configurations.First();
        config.WindowPositionX = windowPosition[0];
        config.WindowPositionY = windowPosition[1];

        await _dbContext.SaveChangesAsync();
    }
    
    private async void UpdateWindowSize()
    {
        var mainWindow = Electron.WindowManager.BrowserWindows.First();
        var windowSize = await mainWindow.GetSizeAsync();

        var config = _dbContext.Configurations.First();
        config.WindowWidth = windowSize[0];
        config.WindowHeight = windowSize[1];

        await _dbContext.SaveChangesAsync();
    }
}