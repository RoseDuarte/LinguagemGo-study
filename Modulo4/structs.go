package main

import "fmt"

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipos diferentes

// type <nome da nossa estrutura> struct {<campos>}
type Pessoa struct {
	Nome string
	Idade int
}

func main() {
	fmt.Println(Pessoa{"Maria", 28})
	fmt.Println(Pessoa{Nome: "João", Idade: 4})
	fmt.Println(Pessoa{Nome:"Maria"})

	p1 := Pessoa{Nome: "Bob", Idade: 2}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	p1.Idade = 3
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Patrick", Idade:2}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)
}

