using ElectronNET.API;
using ElectronNET.API.Entities;

namespace FleetDashClient;

public class ElectronBootstrap
{
    public static async void Bootstrap()
    {
        var browserWindow = await Electron.WindowManager.CreateWindowAsync(new BrowserWindowOptions
        {
            Width = 1152,
            Height = 940,
            Show = false
        });
        await browserWindow.WebContents.Session.ClearCacheAsync();
        browserWindow.OnReadyToShow += () => browserWindow.Show();
        browserWindow.SetTitle("JetBrains!");
    }
}