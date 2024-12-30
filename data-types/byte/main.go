package main

import (
	"fmt"
)

func main() {
	// Original string
	original := "Hello, Go!"

	// Convert the string to a slice of bytes
	bytes := []byte(original)
	fmt.Printf("Original bytes: %v\n", bytes)

	// Modify the bytes (shift each character to the next in the ASCII table)
	for i := range bytes {
		bytes[i] += 1
	}

	// Display the modified bytes
	fmt.Printf("Modified bytes: %v\n", bytes)

	// Convert the bytes back to a string
	modified := string(bytes)
	fmt.Printf("Modified string: %s\n", modified)
}
