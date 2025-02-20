package main

import "fmt"

func main() {
	// Declaração de uma variável comum
	var num int = 10

	// Criando um ponteiro que armazenará o endereço de memória de `num`
	var ptr *int = &num

	// Exibindo o valor da variável e seu endereço na memória
	fmt.Println("Valor de num:", num)                  // 10
	fmt.Println("Endereço de num:", &num)              // Exemplo: 0xc0000140a8
	fmt.Println("Valor armazenado no ponteiro:", ptr)  // Mesmo endereço de num
	fmt.Println("Valor apontado pelo ponteiro:", *ptr) // 10 (desreferência)

	// Modificando o valor da variável através do ponteiro
	*ptr = 20
	fmt.Println("Novo valor de num após alteração via ponteiro:", num) // 20

	// Exemplo com função que modifica uma variável via ponteiro
	changeValue(&num)
	fmt.Println("Valor de num após a função:", num) // 100
}

// Função que recebe um ponteiro e altera o valor original da variável
func changeValue(n *int) {
	*n = 100 // Modifica diretamente o valor da variável original
}
