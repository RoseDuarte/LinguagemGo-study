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

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Maria", 28})
	fmt.Println(Pessoa{Nome: "João", Idade: 4})
	fmt.Println(Pessoa{Nome:"Maria"})

	p1 := Pessoa{Nome: "Bob", Idade: 2}
	//fmt.Println(p1.Nome)
	//fmt.Println(p1.Idade)

	//p1.Idade = 3
	//fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Patrick", Idade: 2}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	//fmt.Println(pessoas)

	//structs + map
	// alunos := map[string][]Pessoa{}
	// alunos["programação"] = pessoas
	// fmt.Println(alunos)

	// var alunos = map[string][]Pessoa {
	// 	"Programação": {{Nome: "Maria", Idade: 28}, {Nome: "João", Idade: 4}},
	// 	"Engenharia": {{Nome: "Maria2", Idade: 28}, {Nome: "João2", Idade: 4}},
	// }
	// fmt.Println(alunos)

	// Struct herdando campos de outra struct
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)

	
}

