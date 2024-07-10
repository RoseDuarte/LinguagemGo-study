package main

import "fmt"

// Ponteiros: um ponteiro nada mais é do que uma variável que ao invés
// de armazenar um valor (1, "olá", true...), armazena um endereço de memória

func main() {
	// Memória -> essa memória tem um endereço -> esse endereço guarda um valor

	i := 1
	fmt.Println("Valor inicial:", i)
}