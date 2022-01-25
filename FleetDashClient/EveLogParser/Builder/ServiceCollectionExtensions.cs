using Microsoft.Extensions.DependencyInjection;

namespace EveLogParser.Builder;

public static class ServiceCollectionExtensions
{
    public static IServiceCollection AddEveLogParser(this IServiceCollection services, Action<EveLogParserOptions> options)
    {
        var builder = new EveLogParserOptions();
        options(builder);
        return services;
    }
}