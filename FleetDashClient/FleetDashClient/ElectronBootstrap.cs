using ElectronNET.API;
using ElectronNET.API.Entities;
using FleetDashClient.Models;

namespace FleetDashClient;

public class ElectronBootstrap
{
    
    public static async void Bootstrap(Configuration config)
    {
        var browserWindow = await Electron.WindowManager.CreateWindowAsync(new BrowserWindowOptions
        {
            Width = config.WindowWidth,
            Height = config.WindowHeight,
            X = config.WindowPositionX,
            Y = config.WindowPositionY,
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