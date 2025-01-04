package main

import "fmt"

type User struct {
	Name     string
	Age      int
	LastName string
	Gender   string
	Address  string
}

func main() {

	user := User{
		Name:     "Victor",
		Age:      28,
		LastName: "Lira",
		Gender:   "M",
		Address:  "Jardim botânico, Clevert - SP",
	}

	fmt.Printf(" FirstName: %s! \n LastName: %s \n Age: %d \n Gender: %s \n %s", user.Name, user.LastName, user.Age, user.Gender, user.Address)
}

// A struct em Go se assemelha ao type ou interface em TypeScript.
// A utilidade de uma struct em Go (como a User no seu exemplo) é definir um tipo de dado composto que pode armazenar
// diferentes informações relacionadas. As structs permitem agrupar variáveis (campos) de tipos diferentes sob um único nome,
//o que facilita a organização, manipulação e passagem de dados entre funções. No seu caso, a struct User armazena informações sobre um
//usuário, como nome, sobrenome e idade.
