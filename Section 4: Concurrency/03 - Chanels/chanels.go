package main

import (
	"fmt"
	"time"
)

// Os channels são utilizados para comunicação segura entre goroutines. Eles permitem que goroutines
// compartilhem dados sem a necessidade de locks (mutexes), garantindo sincronização automática.
// Em outras palavras, um channel permite que uma goroutine envie dados e outra os receba de forma sincronizada.

func main() {
	// Criando um canal de string
	messageChannel := make(chan string)

	// Goroutine que envia dados para o canal
	go func() {
		fmt.Println("Enviando mensagem para o canal...")
		time.Sleep(2 * time.Second)       // Simula um atraso antes de enviar
		messageChannel <- "Olá, Channel!" // Envia um valor para o canal
		fmt.Println("Mensagem enviada!")
	}()

	// Recebendo o valor do canal (bloqueia até que um valor seja recebido)
	fmt.Println("Aguardando mensagem...")
	message := <-messageChannel // Aguarda a goroutine enviar o valor
	fmt.Println("Mensagem recebida:", message)

	fmt.Println("Programa finalizado.")
}
