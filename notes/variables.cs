// Numéricos Inteiros
byte idade = 25;         // 0 a 255
short temperatura = -15; // -32.768 a 32.767
int populacao = 100000;  // -2.147.483.648 a 2.147.483.647
long distancia = 9223372036854775807L; // 64 bits


// Numéricos de Ponto Flutuante
float peso = 65.5f;      // Suporta precisão de 7 dígitos
double altura = 1.75;    // Suporta precisão de 15-16 dígitos
decimal salario = 1200.75m; // Alta precisão, ideal para valores financeiros


// Caractere
char letra = 'A';        // Um único caractere
char simbolo = '😊';      // Caracter Unicode


// Lógicos
bool estaLigado = true;  // Verdadeiro ou falso


// struct
struct Ponto
{
    public int X;
    public int Y;
}


Ponto ponto = new Ponto { X = 5, Y = 10 };
Console.WriteLine($"Ponto: ({ponto.X}, {ponto.Y})");

// enum:
enum DiaDaSemana { Segunda, Terca, Quarta, Quinta, Sexta, Sabado, Domingo }
DiaDaSemana hoje = DiaDaSemana.Sexta;
Console.WriteLine($"Hoje é: {hoje}");


// 2. ----- Tipos de Referência -----
string mensagem = "Olá, Mundo!";
Console.WriteLine(mensagem);


// Object
object valor = 42;      // Pode armazenar qualquer tipo
valor = "Texto";        // Agora é uma string
Console.WriteLine(valor);


// Classe
class Pessoa
{
    public string Nome { get; set; }
    public int Idade { get; set; }
}
Pessoa pessoa = new Pessoa { Nome = "Victor", Idade = 30 };
Console.WriteLine($"Nome: {pessoa.Nome}, Idade: {pessoa.Idade}");


// Interface
interface IAnimal
{
    void FazerSom();
}

class Cachorro : IAnimal
{
    public void FazerSom() => Console.WriteLine("Au Au!");
}

IAnimal cachorro = new Cachorro();
cachorro.FazerSom();


// Delegate
delegate void MostrarMensagem(string mensagem);

void Exibir(string msg) => Console.WriteLine(msg);

MostrarMensagem mostrar = Exibir;
mostrar("Olá, Delegates!");


// Array
int[] numeros = { 1, 2, 3, 4, 5 };
Console.WriteLine($"Primeiro número: {numeros[0]}");


// Dynamic
dynamic valorDinamico = 10;
valorDinamico = "Texto Dinâmico";
Console.WriteLine(valorDinamico);


// 3. Tipos Nulos
int? idade = null;
if (idade.HasValue)
    Console.WriteLine($"Idade: {idade}");
else
    Console.WriteLine("Idade não definida.");


// 4. Tipos Anônimos
var produto = new { Nome = "Laptop", Preco = 3500.50 };
Console.WriteLine($"Produto: {produto.Nome}, Preço: {produto.Preco}");


// 5. Tipos Implícitos
var nome = "Victor";     // Compilador infere que é string
var idade = 30;          // Compilador infere que é int
Console.WriteLine($"Nome: {nome}, Idade: {idade}");


// 6. Tipos Dinâmicos
dynamic dinamico = 10;
Console.WriteLine($"Número: {dinamico}");
dinamico = "Agora sou texto";
Console.WriteLine($"Texto: {dinamico}");


// 7. Tipos de Ponteiro (Unsafe Context)
// Requer habilitar código inseguro (unsafe) no projeto.
unsafe
{
    int numero = 10;
    int* ptr = &numero;
    Console.WriteLine($"Valor: {numero}, Endereço: {(int)ptr}");
}

