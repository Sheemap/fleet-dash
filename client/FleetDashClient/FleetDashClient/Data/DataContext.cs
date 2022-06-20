using FleetDashClient.Models;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;

namespace FleetDashClient.Data;

public class DataContext : DbContext
{
    public DataContext(DbContextOptions<DataContext> options) : base(options)
    {
    }
    
    public DbSet<Character> Characters { get; set; }
    public DbSet<Token> Tokens { get; set; }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Character>(x =>
            x.HasKey(y => y.Id));
        modelBuilder.Entity<Character>(x =>
            x.HasMany(y => y.Tokens)
                .WithOne(y => y.Character)
                .HasForeignKey(y => y.CharacterId));

        
        modelBuilder.Entity<Token>(x =>
            x.HasKey(y => new { y.CharacterId, y.AccessToken }));
        modelBuilder.Entity<Token>(x =>
            x.HasIndex(y => y.ExpiresAt));

        base.OnModelCreating(modelBuilder);
        
        if (Database.ProviderName == "Microsoft.EntityFrameworkCore.Sqlite")
        {
            // SQLite does not have proper support for DateTimeOffset via Entity Framework Core, see the limitations
            // here: https://docs.microsoft.com/en-us/ef/core/providers/sqlite/limitations#query-limitations
            // To work around this, when the Sqlite database provider is used, all model properties of type DateTimeOffset
            // use the DateTimeOffsetToBinaryConverter
            // Based on: https://github.com/aspnet/EntityFrameworkCore/issues/10784#issuecomment-415769754
            // This only supports millisecond precision, but should be sufficient for most use cases.
            foreach (var entityType in modelBuilder.Model.GetEntityTypes())
            {
                var properties = entityType.ClrType.GetProperties().Where(p => p.PropertyType == typeof(DateTimeOffset)
                                                                               || p.PropertyType == typeof(DateTimeOffset?));
                foreach (var property in properties)
                {
                    modelBuilder
                        .Entity(entityType.Name)
                        .Property(property.Name)
                        .HasConversion(new DateTimeOffsetToBinaryConverter());
                }
            }
        }
    }
}