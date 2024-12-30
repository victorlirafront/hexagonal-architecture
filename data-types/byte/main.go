package main

import (
	"fmt"
)

func main() {
	// String original
	original := "Hello, Go!"

	// Converte a string para um slice de bytes
	bytes := []byte(original)
	fmt.Printf("Bytes originais: %v\n", bytes)

	// Modifica os bytes (desloca cada caractere para o próximo na tabela ASCII)
	for i := range bytes {
		bytes[i] += 1
	}

	// Exibe os bytes modificados
	fmt.Printf("Bytes modificados: %v\n", bytes)

	// Converte os bytes de volta para uma string
	modified := string(bytes)
	fmt.Printf("String modificada: %s\n", modified)
}

// O tipo byte em Go é um alias para o tipo uint8, representando um valor numérico de 8 bits sem sinal, variando de 0 a 255. Ele é amplamente utilizado para manipular dados binários, como caracteres individuais, arquivos e fluxos de rede.

// Principais Características:
// Interpretação Semântica:

// Usado para indicar que o dado representa um byte de informação, em vez de um número genérico.
// Manipulação de Strings:

// Strings em Go são sequências imutáveis de bytes. O tipo byte permite acessar e modificar esses valores.
// Facilidade em Operações Binárias:

// Ideal para processar dados binários ou converter informações entre formatos.
