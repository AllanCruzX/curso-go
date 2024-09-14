package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
}

type Cliente struct {
	Nome        string
	Idade       int
	Ativo       bool
	EnderecoObj Endereco
}

func main() {

	allan := Cliente{
		Nome:  "Allan",
		Idade: 36,
		Ativo: true,
	}

	allan.EnderecoObj.Cidade = "São Paulo"
	allan.EnderecoObj.Estado = "SP"
	allan.EnderecoObj.Numero = 10
	allan.EnderecoObj.Logadouro = "Rua do Endereço que é um logadouro"
	fmt.Println(allan)

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", allan.Nome, allan.Idade, allan.Ativo)

}
