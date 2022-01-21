using ElectronNET.API;
using FleetDashClient.Data;
using FleetDashClient.Models.Events;
using FleetDashClient.ViewModels;

namespace FleetDashClient.Services;

public class WorkerService : BackgroundService
{
    private ILogParserService LogParserService { get; set; }
    private ILogReaderService LogReaderService { get; set; }
    
    private readonly IServiceProvider _serviceProvider;

    private bool _handlersInitialized;
    
    private DataContext _dbContext;
    private ICharacterService _characterService;
    private LogViewModel _logViewModel;

    public WorkerService(IServiceProvider services)
    {
        _serviceProvider = services;
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
        while (!stoppingToken.IsCancellationRequested)
        {
            try
            {
                // Hosted services are singleton, but everything else is scoped.
                // So we must make our own scope to get the services we need.
                using var scope = _serviceProvider.CreateScope();
        
                _dbContext = scope.ServiceProvider.GetRequiredService<DataContext>();
                _characterService = scope.ServiceProvider.GetRequiredService<ICharacterService>();
                _logViewModel = scope.ServiceProvider.GetRequiredService<LogViewModel>();
        
                StartLogProcessing();
                
                await EnsureElectronHandlersInitialized(stoppingToken);

                while (!stoppingToken.IsCancellationRequested)
                {
                    Thread.Sleep(1000);
                }
            }
            catch(Exception e)
            {
                Console.WriteLine(e.Message, e.StackTrace);
                if (HybridSupport.IsElectronActive)
                {
                    Electron.Dialog.ShowErrorBox("Unexpected error occured!",
                        $"Something bad happened! We are internally restarting our system, " +
                        $"if this doesnt resolve the issue, please restart the app.\nError:{e.Message}");
                }
            }
        }
    }

    private void StartLogProcessing()
    {
        var config = _dbContext.Configurations.First();
        LogReaderService = new LogReaderService(config.LogDirectory);
        LogParserService = new LogParserService(LogReaderService);
        
        var characters = _dbContext.Characters.ToList();
        var overviewYaml = GetOverviewYaml();
        
        characters.ForEach(x => LogParserService.StartWatchingCharacter(x.Id, overviewYaml));
        LogReaderService.Start();
        
        _characterService.OnCharacterAdded += HandleCharacterAdded;
        _characterService.OnCharacterRemoved += HandleCharacterRemoved;
        _logViewModel.OnLogDirectoryUpdated += HandleLogDirectoryUpdated;
        _logViewModel.OnOverviewFileUpdated += HandleOverviewFileUpdated;
    }
    
    private void HandleOverviewFileUpdated(object sender, PathUpdatedEventArgs overviewPath)
    {
        var watchedChars = LogParserService.WatchedCharacters;
        var overviewYaml = GetOverviewYaml(overviewPath.Path);
        
        watchedChars.ForEach(LogParserService.StopWatchingCharacter);
        watchedChars.ForEach(x => LogParserService.StartWatchingCharacter(x, overviewYaml));
    }

    private void HandleLogDirectoryUpdated(object sender, PathUpdatedEventArgs logDirectory)
    {
        // Changing directory requires a restart of the log reader.
        LogReaderService.Stop();
        StartLogProcessing();
    }
    
    private void HandleCharacterAdded(object sender, CharacterAddedEventArgs args)
    {
        var overviewYaml = GetOverviewYaml();

        LogParserService.StartWatchingCharacter(args.Character.Id, overviewYaml);
    }

    private string? GetOverviewYaml(string? path = null)
    {
        var overviewPath = path ?? _dbContext.Configurations.First().Overview;
        if (string.IsNullOrWhiteSpace(overviewPath)) return null;
        
        using var reader = new StreamReader(overviewPath);
        var overviewYaml = reader.ReadToEnd();

        return overviewYaml;
    } 
    
    private void HandleCharacterRemoved(object sender, CharacterRemovedEventArgs args)
    {
        LogParserService.StopWatchingCharacter(args.Character.Id);
    }

    private async Task EnsureElectronHandlersInitialized(CancellationToken stoppingToken)
    {
        while (!stoppingToken.IsCancellationRequested || !_handlersInitialized)
        {
            AddElectronEventHandlers();
            await Task.Delay(1000, stoppingToken);
        }
    }

    private void AddElectronEventHandlers()
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