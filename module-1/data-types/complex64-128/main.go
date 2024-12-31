package main

import "fmt"

func main() {
	var c1 complex64 = 3 + 4i
	var c2 complex128 = 3 + 4i

	fmt.Println(c1) // 3+4i
	fmt.Println(c2) // 3+4i
}

// Em Go, complex64 e complex128 são tipos numéricos para representar números complexos com precisão de 32 ou 64 bits, respectivamente.
// complex64: representa um número complexo com componentes de 32 bits (float32) para a parte real e imaginária.
// complex128: representa um número complexo com componentes de 64 bits (float64) para a parte real e imaginária.
