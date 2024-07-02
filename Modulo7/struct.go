package main

import "fmt"

type Pessoa struct {
	Nome      string
	Idade     int
	Profissao string
}

func (p Pessoa) ListaNomeEIdade() string {
	return fmt.Sprintf("%s tem %d anos", p.Nome, p.Idade)
}

func main() {
	pessoa := Pessoa {
		Nome: "Maria",
		Idade: 28,
		Profissao: "dev",
	}

	pessoa2 := Pessoa {
		Nome: "Bento",
		Idade: 4,
		Profissao: "dog",
	}

	println(pessoa.ListaNomeEIdade())
	println(pessoa2.ListaNomeEIdade())
}