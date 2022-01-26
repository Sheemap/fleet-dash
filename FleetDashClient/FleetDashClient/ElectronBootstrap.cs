using ElectronNET.API;
using ElectronNET.API.Entities;
using FleetDashClient.Models;

namespace FleetDashClient;

public class ElectronBootstrap
{
    
    public static async void Bootstrap(IConfiguration configuration)
    {
        var config = configuration.Get<Models.Configuration>();
        
        var browserWindow = await Electron.WindowManager.CreateWindowAsync(new BrowserWindowOptions
        {
            Width = config.WindowWidth ?? default,
            Height = config.WindowHeight ?? default,
            X = config.WindowX ?? default,
            Y = config.WindowY ?? default,
            Show = false,
        });
        
        
        var trayItem = new MenuItem
        {
            Label = "Quit",
            Click = () => browserWindow.Close(),
        };

        
        void OnTrayClick(TrayClickEventArgs _, Rectangle __)
        {
            browserWindow.Focus();
        }
        
        Electron.Tray.Show("wwwroote/alert-circle.svg", trayItem);
        Electron.Tray.SetToolTip("FleetDash is running.");
        Electron.Tray.OnClick += OnTrayClick;
        
        await browserWindow.WebContents.Session.ClearCacheAsync();
        browserWindow.OnReadyToShow += () => browserWindow.Show();
        browserWindow.SetTitle("FleetDash");
    }
}