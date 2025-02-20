package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	for num := range ch { // Lê valores do canal
		fmt.Println("Processando:", num)
		time.Sleep(time.Second) // Simula trabalho pesado
	}
}

func main() {
	ch := make(chan int, 3) // Buffer de 3 valores

	go worker(ch) // Inicia a goroutine

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Três tarefas enviadas!")

	time.Sleep(4 * time.Second) // Aguarda tempo suficiente para a goroutine processar
}
