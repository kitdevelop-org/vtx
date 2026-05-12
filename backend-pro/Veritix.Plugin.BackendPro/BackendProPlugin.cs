using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.EntityFrameworkCore;
using Veritix.Plugin.SDK;
using Veritix.Plugin.SDK.Database;
using Veritix.Plugin.BackendPro.Services;
using Veritix.Plugin.BackendPro.Data;

namespace Veritix.Plugin.BackendPro;

public class BackendProPlugin : PluginBase
{
    public override string PluginId => "BackendPro";
    public override string DisplayName => "Backend Pro";
    public override string Version => "1.0.0";
    public override string Description => "Módulo de Backend Pro para Veritix ERP.";
    public override string Author => "FullStackers";
    public override string? Icon => "Box";

    public override void ConfigureServices(IServiceCollection services, IConfiguration configuration)
    {
        // 1. Registrar Contexto de Base de Datos (PostgreSQL)
        var connectionString = configuration.GetConnectionString("DefaultConnection");
        services.AddDbContextFactory<BackendProDbContext>(options =>
            options.UseNpgsql(connectionString));

        // 2. Registrar Migrador
        services.AddScoped<IPluginMigrator, BackendProMigrator>();

        // 3. Registrar Servicios de Negocio
        services.AddScoped<IBackendProService, BackendProService>();
    }

    public override void MapEndpoints(IApplicationBuilder app)
    {
        // El Core de Veritix mapeará automáticamente los controladores 
        // definidos en el namespace de este plugin.
    }

    public override async Task OnInstalledAsync(Guid tenantId, CancellationToken ct = default)
    {
        // Ejecutar migraciones automáticamente al instalar
        // Esto creará el esquema del plugin para el tenant
        await base.OnInstalledAsync(tenantId, ct);
    }
}

