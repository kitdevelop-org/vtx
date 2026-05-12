using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Veritix.Plugin.SDK;

namespace Veritix.Plugin.FrontendPro;

public class FrontendProPlugin : PluginBase
{
    public override string PluginId => "FrontendPro";
    public override string DisplayName => "Frontend Pro";
    public override string Version => "1.0.0";
    public override string Description => "Módulo de Frontend Pro para Veritix ERP.";
    public override string Author => "WebDevs";

    public override void ConfigureServices(IServiceCollection services, IConfiguration configuration)
    {
        // Registrar servicios específicos del plugin
    }

    public override void MapEndpoints(IApplicationBuilder app)
    {
        // Mapear rutas de API (Minimal APIs o controladores)
    }

    public override Task OnInstalledAsync(Guid tenantId, CancellationToken ct = default)
    {
        // Lógica de instalación (migraciones, seed data)
        return base.OnInstalledAsync(tenantId, ct);
    }
}
