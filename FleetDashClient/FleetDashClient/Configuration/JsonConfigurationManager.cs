using Microsoft.Extensions.Options;
using Newtonsoft.Json;
using Serilog;

namespace FleetDashClient.Configuration;

public class JsonConfigurationManager
{
    private readonly ConfigurationOptions _options;
    
    public JsonConfigurationManager(IOptions<ConfigurationOptions> options)
    {
        _options = options.Value;
    }
    
    /// <summary>
    /// Partially updates the config file with the given values.
    /// Only the values that are not null will be updated.
    /// </summary>
    /// <param name="path"></param>
    /// <param name="config"></param>
    public void UpdateConfiguration(Models.Configuration config)
    {
        var currentConfig = JsonConvert.DeserializeObject<Models.Configuration>(File.ReadAllText(_options.Path));
        Log.Debug("Current config: {@Config}", currentConfig);
        var updatedConfig = new Models.Configuration
        {
            LogDirectory = config.LogDirectory ?? currentConfig.LogDirectory,
            OverviewPath = config.OverviewPath ?? currentConfig.OverviewPath,
            WindowHeight = config.WindowHeight ?? currentConfig.WindowHeight,
            WindowWidth = config.WindowWidth ?? currentConfig.WindowWidth,
            WindowX = config.WindowX ?? currentConfig.WindowX,
            WindowY = config.WindowY ?? currentConfig.WindowY
        };
        Log.Debug("Updated config: {@Config}", currentConfig);
        File.WriteAllText(_options.Path, JsonConvert.SerializeObject(updatedConfig, Formatting.Indented));
    }

    public static void EnsureConfigFileExists(string? path)
    {
        if (File.Exists(path)) return;
        
        var config = new Models.Configuration
        {
            WindowHeight = 800,
            WindowWidth = 600,
            WindowX = 0,
            WindowY = 0,
            LogDirectory = null,
            OverviewPath = null,
        };
        File.WriteAllText(path, JsonConvert.SerializeObject(config, Formatting.Indented));
    }
}