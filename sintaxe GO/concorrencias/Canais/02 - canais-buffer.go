package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) // O 2 seria o buffer, ou seja receberia 2 valores no canal mas travaria com o terceiro.
	canal <- "Olá Mundo"

	message := <-canal

	fmt.Println(message)
}
