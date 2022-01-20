using FleetDashClient.Models;
using Microsoft.EntityFrameworkCore;

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
    }
}