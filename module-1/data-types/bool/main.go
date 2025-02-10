package main

import (
	"fmt"
	"math/cmplx"
)

// Declaration of global variables with explicit initialization
var (
	ToBe   bool       = false                // Boolean variable initialized to false
	MaxInt uint64     = 1<<64 - 1            // Maximum value for a 64-bit unsigned integer
	z      complex128 = cmplx.Sqrt(-5 + 12i) // Complex number calculated as the square root of (-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
