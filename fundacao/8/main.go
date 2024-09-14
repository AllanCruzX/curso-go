package main

import (
	"errors"
	"fmt"
)

func main() {

	soma, err := soma(50, 1)

	if err != nil {
		fmt.Printf("Algo deu errado: \"%v\", valor da soma: %d\n", err, soma)
		return
	}
	fmt.Printf("Valor da soma: %d\n", soma)

}

func soma(a, b int) (int, error) {
	soma := a + b
	if soma >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return soma, nil

}
