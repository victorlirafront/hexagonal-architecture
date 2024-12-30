package main

import "fmt"

func main() {
	var a uint = 42
	var b uint8 = 200
	var c uint16 = 60000
	var d uint32 = 4000000000
	var e uint64 = 18000000000000000000

	fmt.Println(a, b, c, d, e)
}

// Em Go, os tipos uint, uint8, uint16, uint32, e uint64 são usados para representar números inteiros não negativos (sem sinal), com diferentes tamanhos em bits.

// uint: Um tipo inteiro sem sinal cujo tamanho depende da arquitetura do sistema (32 bits em sistemas de 32 bits e 64 bits em sistemas de 64 bits).
// uint8: Representa um número inteiro sem sinal com 8 bits (1 byte), com valores que variam de 0 a 255.
// uint16: Representa um número inteiro sem sinal com 16 bits (2 bytes), com valores que variam de 0 a 65.535.
// uint32: Representa um número inteiro sem sinal com 32 bits (4 bytes), com valores que variam de 0 a 4.294.967.295.
// uint64: Representa um número inteiro sem sinal com 64 bits (8 bytes), com valores que variam de 0 a 18.446.744.073.709.551.615.
// Esses tipos são úteis quando você sabe que os números não serão negativos, permitindo o uso de uma gama maior de valores em comparação aos tipos com sinal (int8, int16, etc.).
