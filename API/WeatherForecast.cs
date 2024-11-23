namespace API;

public class WeatherForecast
{
    // O modificador public indica que essa propriedade pode ser acessada de qualquer lugar no código.
    // Se fosse private, por exemplo, a propriedade só poderia ser acessada dentro da própria classe.
    public DateOnly Date { get; set; }

    // O nome da propriedade é TemperatureC, que é um identificador usado para se referir a essa informação na classe ou instância.
    public int TemperatureC { get; set; }

    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);

    public string? Summary { get; set; }
}
