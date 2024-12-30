package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 42
	var ptr *int = &x

	// Convertendo o ponteiro para uintptr
	var addr uintptr = uintptr(unsafe.Pointer(ptr))

	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x (como uintptr):", addr)
}

// Em Go, o tipo uintptr é um tipo inteiro sem sinal usado para armazenar o valor de um ponteiro. Ele é utilizado principalmente para operações de baixo nível, como manipulação de ponteiros em conjunto com funções de C ou operações específicas de memória.
// uintptr: É um tipo inteiro sem sinal que tem o mesmo tamanho que um ponteiro no sistema (ou seja, o tamanho de um ponteiro depende da arquitetura do sistema, geralmente 32 ou 64 bits).
// Uso comum: O uintptr pode ser usado para realizar conversões entre ponteiros e números inteiros, o que pode ser útil em operações que exigem manipulação direta de endereços de memória.

// Para usar uintptr com ponteiros, você precisa importar o pacote unsafe para realizar a conversão, como mostrado no exemplo. O uso de uintptr é recomendado apenas em casos específicos e em baixo nível, pois pode resultar em comportamentos inesperados se mal utilizado.
