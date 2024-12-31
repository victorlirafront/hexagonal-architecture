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

	fmt.Printf("FirstName: %s! LastName: %s Age: %d \n", user.Name, user.LastName, user.Age)
}

// A struct em Go se assemelha ao type ou interface em TypeScript.
// A utilidade de uma struct em Go (como a User no seu exemplo) é definir um tipo de dado composto que pode armazenar
// diferentes informações relacionadas. As structs permitem agrupar variáveis (campos) de tipos diferentes sob um único nome,
//o que facilita a organização, manipulação e passagem de dados entre funções. No seu caso, a struct User armazena informações sobre um
//usuário, como nome, sobrenome e idade.
