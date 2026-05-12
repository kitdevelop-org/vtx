using Microsoft.EntityFrameworkCore;
using Veritix.Plugin.SDK.Database;

namespace Veritix.Plugin.BackendPro.Data;

public class BackendProDbContext(DbContextOptions<BackendProDbContext> options) 
    : DbContext(options), IPluginDbContext
{
    // Define tus DbSets aquí
    // public DbSet<MyEntity> Entities => Set<MyEntity>();

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);
        // Configuración de Fluent API
    }
}
