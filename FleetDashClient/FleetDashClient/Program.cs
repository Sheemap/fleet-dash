using ElectronNET.API;
using FleetDashClient;
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

var dbFileLocation = Path.Combine(location, "fleetdash.db");

builder.Services.AddDbContext<DataContext>(options =>
    options.UseSqlite($"Data Source={dbFileLocation}"));

builder.Services.AddScoped<ICharacterService, CharacterService>();
builder.Services.AddScoped<IndexViewModel>();


var app = builder.Build();

// Ensure DB schema is up to date
using (var scope = app.Services.CreateScope())
{
    var db = scope.ServiceProvider.GetRequiredService<DataContext>();
    db.Database.Migrate();
}

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

if (HybridSupport.IsElectronActive)  ElectronBootstrap.Bootstrap();

app.Run();