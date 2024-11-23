using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

// Marca a classe como um controlador de API e define a rota base para o controlador
[ApiController]
[Route("[controller]")]
public class WeatherForecastController : ControllerBase
{
    // Array de descrições climáticas estáticas para uso aleatório
    private static readonly string[] Summaries = new[]
    {
        "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
    };

    // Campo para o logger injetado, usado para registrar logs
    private readonly ILogger<WeatherForecastController> _logger;

    // Construtor que recebe o logger via injeção de dependência
    public WeatherForecastController(ILogger<WeatherForecastController> logger)
    {
        _logger = logger;
    }

    // Método GET que retorna uma lista de previsões do tempo
    [HttpGet(Name = "GetWeatherForecast")]
    public IEnumerable<WeatherForecast> Get()
    {
        // Gera 5 previsões de tempo aleatórias
        return Enumerable.Range(1, 5).Select(index => new WeatherForecast
        {
            // Define a data como hoje + número de dias (1 a 5)
            Date = DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
            // Define a temperatura aleatória entre -20 e 55
            TemperatureC = Random.Shared.Next(-20, 55),
            // Escolhe uma descrição climática aleatória do array
            Summary = Summaries[Random.Shared.Next(Summaries.Length)]
        })
        // Converte o resultado para um array antes de retornar
        .ToArray();
    }
}
