using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace EveLogParser.Builder;

public static class ServiceCollectionExtensions
{
    /// <summary>
    /// Add EveLogParser to dependency injection, with configured options.
    /// </summary>
    /// <param name="services"></param>
    /// <param name="options"></param>
    /// <returns></returns>
    /// <exception cref="ArgumentNullException"></exception>
    public static IServiceCollection AddEveLogParser(this IServiceCollection services, Action<EveLogParserOptions> options)
    {
        if (options == null) throw new ArgumentNullException(nameof(options));
        services.AddOptions<EveLogParserOptions>().Configure(options);
        RegisterDi(services);
        return services;
    }
    
    /// <summary>
    /// Add EveLogParser to dependency injection, with given IConfiguration.
    /// Configuration will bind to the EveLogParserOptions model, ensure the configuration contains all needed option values
    /// </summary>
    /// <param name="services"></param>
    /// <param name="config"></param>
    /// <returns></returns>
    /// <exception cref="ArgumentNullException"></exception>
    public static IServiceCollection AddEveLogParser(this IServiceCollection services, IConfiguration config)
    {
        if (config == null) throw new ArgumentNullException(nameof(config));
        services.Configure<EveLogParserOptions>(config);
        RegisterDi(services);
        return services;
    }

    private static void RegisterDi(IServiceCollection services)
    {
        services.AddScoped<ILogReaderService, LogReaderService>();
        services.AddScoped<ILogParserService, LogParserService>();
    }
}