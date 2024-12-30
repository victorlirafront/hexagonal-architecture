package main

import "fmt"

// Person is a structure that represents a person.
type Person struct {
	Name   string  // Name of the person
	Age    int     // Age of the person
	Height float64 // Height of the person in meters
}

// Greeting receives a Person and returns a personalized greeting.
func Greeting(p Person) string {
	return "Hello, " + p.Name + "! You are " + fmt.Sprint(p.Age) + " years old and your height is " + fmt.Sprint(p.Height) + " meters."
}

func main() {
	// Variable of type int
	age := 25

	// Variable of type bool
	active := true

	// Variable of type string
	name := "Victor"

	// Variable of type float64
	height := 1.75

	// Using the Person struct
	person := Person{
		Name:   name,
		Age:    age,
		Height: height,
	}

	// Displaying greeting
	fmt.Println(Greeting(person))

	// Primitive types
	fmt.Println("Age:", age)
	fmt.Println("Active:", active)
	fmt.Println("Name:", name)
	fmt.Println("Height:", height)

	// Composite types
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Creates a map where the keys are of type string and the values are of type int.
	mapping := map[string]int{
		"a": 1, // The key "a" is associated with the value 1
		"b": 2, // The key "b" is associated with the value 2
	}

	// Prints the content of the map. Go will display the keys and their respective values.
	fmt.Println("Map:", mapping) // Output: Map: map[a:1 b:2]
}
