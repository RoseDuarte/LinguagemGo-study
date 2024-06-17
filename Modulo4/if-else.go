package main

import "fmt"

// IF / ELSE
// SE / SENÃO

func main() {
	// valor := 3
	// // if (condição) {<ação>}
	// if valor == 1 { //true
	// 	fmt.Println("Valor é igual a 1")
	// } else {
	// 	fmt.Println("Valor não é igual a 1")
	// }

	// if valor == 1 { // true
	// 	fmt.Println("Valor é igual a 1")
	// } else if valor == 2 {
	// 	fmt.Println("Valor é igual a 2")
	// } else {
	// 	fmt.Println("Valor é diferente de 1 e 2")
	// }

	numero := 8
	if numero % 2 == 0 {
		fmt.Printf("%d é par", numero)
	} else {
		fmt.Println("%d é impar", numero)
	}
}

// Operações
// Soma: 1 + 1
//Subtração: 2 -1 
// Divisão: 10 / 2
// Multiplicação: 2 * 2
// Resto da divisão por x: 7 % 2 (resto da divisão por 2)


