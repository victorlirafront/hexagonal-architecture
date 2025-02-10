package main

import "fmt"

func main() {
	var a float32 = 3.14
	var b float64 = 3.14159265359

	fmt.Println("float32:", a)
	fmt.Println("float64:", b)
}

// No Go (Golang), float32 e float64 são tipos numéricos usados para representar números de ponto flutuante.
// float32: Representa um número de ponto flutuante com 32 bits de precisão. A faixa de valores que ele pode armazenar é mais limitada, mas é mais eficiente em termos de memória.
// float64: Representa um número de ponto flutuante com 64 bits de precisão. Ele tem uma faixa de valores maior e é mais preciso, sendo o tipo de ponto flutuante padrão em Go.
