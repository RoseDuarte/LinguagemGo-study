package main

import (
	"fmt"
)

// Loops
// Laços de repetição
// Repetir tarefas

func main () {
	
	// sum := 0

	// // i++ -> soma 1
	// // i-- -> subtrai 1

	// // É como se ao final do loop fosse adicionado 1 ao indice i
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Sum:", sum)
	// 	fmt.Println("Indice:", i)
	// 	sum += i
	// 	// é a mesma coisa que: sum = sum + i
	// 	// sum -= i -> sum = sum - i
	// }
	// fmt.Println(sum)

	// Loop infinito
	// for {
	// 	fmt.Println("Golang do zero")
	// 	time.Sleep(2 * time.Second)
	// }

	// For range
	frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}
	for _, fruta := range frutas {
		fmt.Println("Fruta", fruta)
	}
}