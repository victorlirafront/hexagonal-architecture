package main

import "fmt"

func main() {
	var a int = 42
	var b int8 = 120
	var c int16 = 32000
	var d int32 = 1000000
	var e int64 = 1000000000

	fmt.Println(a, b, c, d, e)
}

// Em Go, os tipos int, int8, int16, int32, e int64 são usados para representar números inteiros com diferentes tamanhos em bits.

// int: É um tipo inteiro cujo tamanho depende da arquitetura do sistema (32 bits em sistemas de 32 bits e 64 bits em sistemas de 64 bits).
// int8: Representa um número inteiro com 8 bits (1 byte), com valores que variam de -128 a 127.
// int16: Representa um número inteiro com 16 bits (2 bytes), com valores que variam de -32.768 a 32.767.
// int32: Representa um número inteiro com 32 bits (4 bytes), com valores que variam de -2.147.483.648 a 2.147.483.647.
// int64: Representa um número inteiro com 64 bits (8 bytes), com valores que variam de -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807.
