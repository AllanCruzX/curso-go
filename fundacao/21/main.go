package main

import (
	"curso-go/21/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	resultado := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v", resultado)
	fmt.Println("Meu UUID", uuid.New())
}
