package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for i := 0; i < 5; i++ { // Repete 5 vezes
		fmt.Println(text)       // Imprime o texto recebido como argumento
		time.Sleep(time.Second) // Pausa por 1 segundo para simular um processo demorado
	}
}

func main() {
	var waitGroup sync.WaitGroup // Cria um WaitGroup para sincronizar as goroutines
	waitGroup.Add(2)             // Define que existem 2 goroutines que precisam ser aguardadas

	// Primeira goroutine anônima
	go func() {
		write("Hello World") // Chama a função write com "Hello World"
		waitGroup.Done()     // Indica que essa goroutine terminou
	}()

	// Segunda goroutine anônima
	go func() {
		write("Hello World 2") // Chama a função write com "Hello World 2"
		waitGroup.Done()       // Indica que essa goroutine terminou
	}()

	waitGroup.Wait() // Aguarda até que as 2 goroutines terminem antes de encerrar `main`

	/*
		go write("Hello World") // Essa linha foi comentada, mas se ativada, criaria uma goroutine concorrente
		write("Programing in Go") // Essa linha executaria normalmente e bloquearia o restante do código
	*/
}
