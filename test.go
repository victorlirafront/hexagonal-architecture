package main

import "fmt"

type User struct {
	Name     string
	Age      int
	LastName string
}

func main() {

	user := User{
		Name:     "Victor",
		Age:      28,
		LastName: "Lira",
	}

	fmt.Printf("Hello, %s!\n", user.Name)
}
