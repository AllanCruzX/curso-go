package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	allan := Cliente{
		Nome:  "Allan",
		Idade: 36,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", allan.Nome, allan.Idade, allan.Ativo)

}
