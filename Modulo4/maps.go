package main

import "fmt"

func main() {
	// idade := map[string]int{}
	// idade["Maria"] = 28
	// idade["João"] = 4
	// fmt.Println(idade)
	// fmt.Println(idade["Maria"])
	// fmt.Println(idade["João"])

	anoNasc := map[string]int{
		"Maria": 1995,
		"João": 2008,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Maria"])
	fmt.Println(anoNasc["João"])
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)	
}