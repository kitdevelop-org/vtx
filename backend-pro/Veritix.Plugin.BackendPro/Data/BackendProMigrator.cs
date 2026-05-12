using Microsoft.EntityFrameworkCore;
using Veritix.Plugin.SDK.Database;

namespace Veritix.Plugin.BackendPro.Data;

public class BackendProMigrator(IDbContextFactory<BackendProDbContext> contextFactory) 
    : PluginMigratorBase<BackendProDbContext>(contextFactory)
{
    // Esta clase se encarga de ejecutar las migraciones del plugin 
    // en el esquema correcto del tenant.
}
