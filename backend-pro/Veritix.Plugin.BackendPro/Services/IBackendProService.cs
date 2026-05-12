namespace Veritix.Plugin.BackendPro.Services;

public interface IBackendProService
{
    string GetHello();
}

public class BackendProService : IBackendProService
{
    public string GetHello() => "Hola desde el servicio del plugin Backend Pro";
}
