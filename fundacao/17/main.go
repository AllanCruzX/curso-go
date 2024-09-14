package main

import "fmt"

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c Conta) simular(valor int) int {
	println("Simulando")
	c.saldo += valor
	return c.saldo
}

func (c *Conta) contemplarSimulacao(valorSimulacao int) {
	println("Aplicando emprestimo")
	c.saldo = valorSimulacao
}

func initConta() {

	conta := Conta{saldo: 0}
	fmt.Printf("O saldo de sua conta é %d\n", conta.saldo)
	valorSimulacao := conta.simular(20000)
	fmt.Printf("O valor ds simulação é %d\n", valorSimulacao)
	fmt.Printf("O saldo de sua conta é %d\n", conta.saldo)
	conta.contemplarSimulacao(valorSimulacao)
	fmt.Printf("O saldo de sua conta é %.d\n", conta.saldo)

	fmt.Printf("---------------------\n")

	contaX := NewConta()
	fmt.Printf("O saldo de sua conta é %d\n", contaX.saldo)
	valorSimulacaoX := contaX.simular(20000)
	fmt.Printf("O valor ds simulação é %d\n", valorSimulacaoX)
	fmt.Printf("O saldo de sua conta é %d\n", contaX.saldo)
	contaX.contemplarSimulacao(valorSimulacaoX)
	fmt.Printf("O saldo de sua conta é %.d\n", contaX.saldo)

}

func main() {
	initConta()
}
