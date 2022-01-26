using ElectronNET.API;
using EveLogParser.Builder;
using FleetDashClient;
using FleetDashClient.Configuration;
using FleetDashClient.Data;
using FleetDashClient.Services;
using FleetDashClient.ViewModels;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

builder.WebHost.UseElectron(args);

// Add services to the container.
builder.Services.AddRazorPages();
builder.Services.AddServerSideBlazor();


var location = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData), "FleetDash");
if (!Directory.Exists(location))
{
    Directory.CreateDirectory(location);
}

var configFileLocation = Path.Combine(location, "config.json");
JsonConfigurationManager.EnsureConfigFileExists(configFileLocation);
builder.Configuration.AddJsonFile(configFileLocation, false, true);

var dbFileLocation = Path.Combine(location, "fleetdash.db");

builder.Services.AddDbContext<DataContext>(options =>
    options.UseSqlite($"Data Source={dbFileLocation}"));

builder.Services.AddEveLogParser(builder.Configuration);

builder.Services.AddScoped<JsonConfigurationManager>();
builder.Services.AddScoped<ICharacterService, CharacterService>();
builder.Services.AddScoped<IndexViewModel>();
builder.Services.AddScoped<LogViewModel>();
builder.Services.AddScoped<EveLoginViewModel>();

builder.Services.AddHostedService<WorkerService>();


var app = builder.Build();

// Ensure DB schema is up to date
using (var scope = app.Services.CreateScope())
{
    var db = scope.ServiceProvider.GetRequiredService<DataContext>();
    db.Database.Migrate();
}

if (HybridSupport.IsElectronActive) ElectronBootstrap.Bootstrap(builder.Configuration);


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