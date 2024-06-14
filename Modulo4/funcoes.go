package main

import "fmt"


func main() {
	somaDosValores := soma(42, 13)
	fmt.Println(somaDosValores)

	sub := subtração(10, 5)
	fmt.Println(sub)

	nome1, nome2 := printaNome("Maria")
	fmt.Println(nome1)
	fmt.Println(nome2)

	nome, sobrenome := printaNomeCompleto("Maria", "Silva")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}

// Função começando com letra minúscula:
// Função é PRIVADA
// A função só pode ser utilizada no próprio pacote 
func printaNome(nome string) (string, string) {
	return nome, nome
}

func printaNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// Função começando com letra maiúscula:
// Função é PÚBLICA
// A função pode ser utilizada fora do próprio pacote
// Como utilizaria fora do pacote: main.PrintaNome(nome)
func PrintaNome(nome string) string {
	return nome
}

func subtração(x int, y int) int {
	return x - y 
}

func soma(x int, y int) int {
	return x + y 
}