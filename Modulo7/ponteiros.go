package main

import "fmt"

// Ponteiros: um ponteiro nada mais é do que uma variável que ao invés
// de armazenar um valor (1, "olá", true...), armazena um endereço de memória

func zeroValue(i int) {
	i = 0
	fmt.Println("End de memória do i dentro da func: ", &i)
}

func zeroPointer(i *int) {
	*i = 0
}

func main() {
	// Memória -> essa memória tem um endereço -> esse endereço guarda um valor

	i := 1
	fmt.Println("Valor inicial:", i)
	fmt.Println("Valor end de memória: ", &i) // &var -> apontando para o end de memória da variável

	// a := &i
	// fmt.Println("Variavel a:", a)
	// fmt.Println("Variavel *a:", *a) // quando usa o *, estamos falando do endereço! E não do valor
	// fmt.Println("Variavel &a:", &a)

	zeroValue(i)
	fmt.Println("zeroval:", i)

	zeroPointer(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	

}