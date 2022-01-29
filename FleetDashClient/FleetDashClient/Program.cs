using ElectronNET.API;
using EveLogParser.Builder;
using FleetDashClient;
using FleetDashClient.Configuration;
using FleetDashClient.Data;
using FleetDashClient.Models;
using FleetDashClient.Protobuf;
using FleetDashClient.Services;
using FleetDashClient.ViewModels;
using Microsoft.EntityFrameworkCore;
using Serilog;
using Serilog.Exceptions;

var appLocation = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData), "FleetDash");
if (!Directory.Exists(appLocation))
{
    Directory.CreateDirectory(appLocation);
}

Log.Logger = new LoggerConfiguration()
    .MinimumLevel.Information()
    .Enrich.WithExceptionDetails()
    .WriteTo.Console()
    .WriteTo.File($"{appLocation}/logs/log.txt", rollingInterval: RollingInterval.Day, retainedFileCountLimit: 7)
    .CreateLogger();

var builder = WebApplication.CreateBuilder(args);

builder.WebHost.UseElectron(args);
builder.WebHost.UseSerilog();

// Add services to the container.
builder.Services.AddRazorPages();
builder.Services.AddServerSideBlazor();


var configFileLocation = Path.Combine(appLocation, "config.json");
JsonConfigurationManager.EnsureConfigFileExists(configFileLocation);
builder.Configuration.AddJsonFile(configFileLocation, false, true);

builder.Services.AddOptions<ConfigurationOptions>().Configure(opts => opts.Path = configFileLocation);
builder.Services.Configure<Configuration>(builder.Configuration);

var dbFileLocation = Path.Combine(appLocation, "fleetdash.db");

builder.Services.AddDbContext<DataContext>(options =>
    options.UseSqlite($"Data Source={dbFileLocation}"));

builder.Services.AddEveLogParser(builder.Configuration);

builder.Services.AddEventAggregator(options =>
    options.AutoRefresh = true);
builder.Services.AddScoped<JsonConfigurationManager>();
builder.Services.AddScoped<ICharacterService, CharacterService>();
builder.Services.AddScoped<IndexViewModel>();
builder.Services.AddScoped<LogViewModel>();
builder.Services.AddScoped<EveLoginViewModel>();
builder.Services.AddScoped<GrpcLogShipper>();

builder.Services.AddGrpcClient<FleetDashService.FleetDashServiceClient>(o =>
{
    o.Address = new Uri("http://localhost:50051");
});

builder.Services.AddHostedService<WorkerService>();


var app = builder.Build();

// Ensure DB schema is up to date
using (var scope = app.Services.CreateScope())
{
    var db = scope.ServiceProvider.GetRequiredService<DataContext>();
    db.Database.Migrate();
}

if (HybridSupport.IsElectronActive) ElectronBootstrap.Bootstrap(builder.Configuration, app);


// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Error");
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

app.UseHttpsRedirection();

app.UseStaticFiles();

app.UseRouting();

app.MapBlazorHub();
app.MapFallbackToPage("/_Host");

app.Run();