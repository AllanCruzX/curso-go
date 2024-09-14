package main

import "fmt"

func main() {
	salarios := map[string]int{"Allan": 10000, "Bia": 2000, "Ivo": 3000}
	fmt.Println(salarios)
	fmt.Println(salarios["Allan"])

	delete(salarios, "Allan")
	fmt.Println(salarios)

	salarios["Allan"] = 230000
	fmt.Println(salarios["Allan"])

	novoSalario := make(map[string]int)
	novoSalario["Allan"] = 500000
	fmt.Println(novoSalario)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é de %d \n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é de %d \n", salario)
	}
}
