package main

import (
	"fmt"
)

func main() {

	// posicao := 3
	// switch posicao {
	// case 1:
	// 	fmt.Println("Primeiro lugar")
	// case 2:
	// 	fmt.Println("Segundo lugar")
	// case 3:
	// 	fmt.Println("Terceiro lugar")
	// }

	nome := "bob"
	switch nome {
	case "maria":
		fmt.Println("É uma pessoa")
	case "Thor":
		fmt.Println("É um cachorro")
	case "bob":
		fmt.Println("É um personagem")

	}
}