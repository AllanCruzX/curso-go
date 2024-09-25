package main

import (
	"fmt"

	"github.com/AllanCruzX/curso-go/packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
}
