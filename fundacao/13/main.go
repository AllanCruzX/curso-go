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

func (cliente Cliente) Desativar() {
	cliente.Ativo = false
	fmt.Printf("O cliene %s foi desativado\n", cliente.Nome)
}

func (cliente Cliente) Ativar() {
	cliente.Ativo = true
	fmt.Printf("O cliene %s foi ativado\n", cliente.Nome)

}

func main() {

	allan := Cliente{
		Nome:  "Allan",
		Idade: 36,
	}

	allan.EnderecoObj.Cidade = "São Paulo"
	allan.EnderecoObj.Estado = "SP"
	allan.EnderecoObj.Numero = 10
	allan.EnderecoObj.Logadouro = "Rua do Endereço que é um logadouro"
	fmt.Println(allan)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", allan.Nome, allan.Idade, allan.Ativo)
	allan.Ativar()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", allan.Nome, allan.Idade, allan.Ativo)

}
