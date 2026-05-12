using Microsoft.AspNetCore.Mvc;
using Veritix.Plugin.SDK;

namespace Veritix.Plugin.BackendPro.Controllers;

[ApiController]
[Route("api/backendpro")]
public class BackendProController : ControllerBase
{
    [HttpGet("status")]
    public IActionResult GetStatus()
    {
        return Ok(new { 
            status = "Operational", 
            plugin = "Backend Pro",
            timestamp = DateTime.UtcNow 
        });
    }
}
