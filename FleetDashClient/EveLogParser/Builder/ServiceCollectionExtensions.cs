using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace EveLogParser.Builder;

public static class ServiceCollectionExtensions
{
    public static IServiceCollection AddEveLogParser(this IServiceCollection services, Action<EveLogParserOptions> options)
    {
        if (options == null) throw new ArgumentNullException(nameof(options));
        var builder = new EveLogParserOptions();
        options(builder);
        return services;
    }
    
    public static IServiceCollection AddEveLogParser(this IServiceCollection services, IConfiguration config)
    {
        if (config == null) throw new ArgumentNullException(nameof(config));
        services.Configure<EveLogParserOptions>(config.Bind);
        return services;
    }
}