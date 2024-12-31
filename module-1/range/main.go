package main

import "fmt"

// Declare a slice containing powers of 2
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// Iterate over the 'pow' slice using range
	// 'i' is the index, and 'v' is the value of the element in the slice
	for index, value := range pow {
		// Print the index as the power and the corresponding value
		fmt.Printf("2**%d = %d\n", index, value)

	}
}
