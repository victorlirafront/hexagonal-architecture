package main

import (
	"fmt"
)

func main() {
	// String containing Unicode characters
	text := "Olá, 🌍!"

	// Converting to a slice of runes
	runes := []rune(text)

	fmt.Println("Individual runes:")
	for i, r := range runes {
		fmt.Printf("Index: %d, Rune: %c, Unicode: U+%04X\n", i, r, r)
	}

	// Modifying a rune
	runes[5] = '👋'

	// Converting back to string
	modifiedText := string(runes)
	fmt.Printf("Modified text: %s\n", modifiedText)
}

// In Go, rune is an alias for the int32 type. It is used to represent Unicode characters, allowing you to store any Unicode code point. Unlike byte (which is an alias for uint8 and typically used to represent raw data), rune deals with characters in a way that supports multiple bytes.
