using System.Runtime.InteropServices;

namespace FleetDashClient;

public static class Utilities
{
    public static string GetDefaultLogDirectory()
    {
        if (RuntimeInformation.IsOSPlatform(OSPlatform.Windows))
        {
            return Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments) + @"\EVE\logs\Gamelogs";
        }
        
        if (RuntimeInformation.IsOSPlatform(OSPlatform.Linux))
        {
            return Directory.Exists("~/Documents/EVE/logs/") ?
                "~/Documents/EVE/logs/Gamelogs" :
                "~/.local/share/Steam/steamapps/compatdata/8500/pfx/drive_c/users/steamuser/My Documents/EVE/logs/Gamelogs";
        }
        
        if (RuntimeInformation.IsOSPlatform(OSPlatform.OSX))
        {
            return Directory.Exists("~/Documents/EVE/logs") ?
                "~/Documents/EVE/logs/Gamelogs" :
                "~/Library/Application Support/EVE Online/p_drive/User/My Documents/EVE/logs/Gamelogs";
        }

        return string.Empty;
    }
}