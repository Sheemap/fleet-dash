using ElectronNET.API;
using ElectronNET.API.Entities;
using FleetDashClient.Models;
using Serilog;

namespace FleetDashClient;

public class ElectronBootstrap
{
    
    public static async void Bootstrap(IConfiguration configuration, WebApplication app)
    {
        var config = configuration.Get<Models.Configuration>();
        
        var browserWindow = await Electron.WindowManager.CreateWindowAsync(new BrowserWindowOptions
        {
            Width = config.WindowWidth ?? 600,
            Height = config.WindowHeight ?? 800,
            X = config.WindowX ?? default,
            Y = config.WindowY ?? default,
            Show = false,
        });
        
        
        var trayItem = new MenuItem
        {
            Label = "Quit",
            Click = () =>
            {
                browserWindow.Close();
                app.StopAsync();
            },
        };
        
        
        void OnTrayClick(TrayClickEventArgs _, Rectangle __)
        {
            browserWindow.Focus();
        }
        
        Electron.Tray.Show("wwwroote/alert-circle.svg", trayItem);
        Electron.Tray.SetToolTip("FleetDash is running.");
        Electron.Tray.OnClick += OnTrayClick;
        
        await browserWindow.WebContents.Session.ClearCacheAsync();
        browserWindow.OnReadyToShow += () =>
        {
            browserWindow.Show();
        };
        browserWindow.SetTitle("FleetDash");
    }
}