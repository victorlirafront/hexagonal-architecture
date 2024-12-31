package main

import "fmt"

func main() {
	fmt.Printf("String: %s, Inteiro: %d, Float: %.2f, Booleano: %t\n", "texto", 42, 3.14, true)
}

// Principais especificadores de formato:
// %s → Representa uma string.
// %d → Representa um número inteiro (decimal).
// %f → Representa um número de ponto flutuante.
// %t → Representa valores booleanos (true ou false).
// %v → Representa o valor no formato padrão (útil para tipos variados).
// %T → Representa o tipo do valor.
