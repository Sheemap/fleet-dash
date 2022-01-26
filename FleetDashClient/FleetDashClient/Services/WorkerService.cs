using ElectronNET.API;
using EveLogParser;
using EveLogParser.Models.Events;
using FleetDashClient.Configuration;
using FleetDashClient.Data;
using FleetDashClient.Models.Events;
using FleetDashClient.ViewModels;

namespace FleetDashClient.Services;

public class WorkerService : BackgroundService
{
    private readonly IServiceProvider _serviceProvider;

    private bool _handlersInitialized;
    
    private DataContext _dbContext;
    private ICharacterService _characterService;
    private LogViewModel _logViewModel;
    private ILogParserService _logParserService;
    private JsonConfigurationManager _configManager;
    private GrpcLogShipper _logShipper;

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
                 _logParserService = scope.ServiceProvider.GetRequiredService<ILogParserService>();
                 _configManager = scope.ServiceProvider.GetRequiredService<JsonConfigurationManager>();
                 _logShipper = scope.ServiceProvider.GetRequiredService<GrpcLogShipper>();
        
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
        var characters = _dbContext.Characters.ToList();
        
        characters.ForEach(x => _logParserService.StartWatchingCharacter(x.Id));
        
        _characterService.OnCharacterAdded += HandleCharacterAdded;
        _characterService.OnCharacterRemoved += HandleCharacterRemoved;
        _logViewModel.OnLogDirectoryUpdated += HandleLogDirectoryUpdated;
        _logViewModel.OnOverviewFileUpdated += HandleOverviewFileUpdated;

        _logParserService.Start();
    }
    
    private void HandleOverviewFileUpdated(object sender, PathUpdatedEventArgs overviewPath)
    {
        var config = new Models.Configuration
        {
            OverviewPath = overviewPath.Path,
        };
        _configManager.UpdateConfiguration(config);
    }

    private void HandleLogDirectoryUpdated(object sender, PathUpdatedEventArgs logDirectory)
    {
        var config = new Models.Configuration
        {
            LogDirectory = logDirectory.Path,
        };
        _configManager.UpdateConfiguration(config);
    }
    
    private void HandleCharacterAdded(object sender, CharacterAddedEventArgs args)
    {
        _logParserService.StartWatchingCharacter(args.Character.Id);
    }

    private void HandleCharacterRemoved(object sender, CharacterRemovedEventArgs args)
    {
        _logParserService.StopWatchingCharacter(args.Character.Id);
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
        
        var config = new Models.Configuration
        {
            WindowX = windowPosition[0],
            WindowY = windowPosition[1],
        };
        _configManager.UpdateConfiguration(config);
    }
    
    private async void UpdateWindowSize()
    {
        var mainWindow = Electron.WindowManager.BrowserWindows.First();
        var windowSize = await mainWindow.GetSizeAsync();
        
        var config = new Models.Configuration
        {
            WindowWidth = windowSize[0],
            WindowHeight = windowSize[1],
        };
        _configManager.UpdateConfiguration(config);
    }
}