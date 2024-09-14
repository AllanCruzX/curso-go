package main

import "fmt"

func SomaInteiro(mapa map[string]int) int {
	var soma int
	for _, valor := range mapa {
		soma += valor
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}

	return soma
}

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T comparable](a, b T) bool {
	return a == b
}
func main() {
	mapa := map[string]int{"Allan": 10000, "Bia": 20000, "Ivo": 1000}
	mapa2 := map[string]float64{"Allan": 10000, "Bia": 20000, "Ivo": 1000}
	mapa3 := map[string]MyNumber{"Allan": 10000, "Bia": 20000, "Ivo": 1000}
	fmt.Printf("SomaInteiro %d\n", SomaInteiro(mapa))
	fmt.Printf("SomaFloat: %.2f\n", SomaFloat(mapa2))

	fmt.Printf("Soma mapa: %d\n", Soma(mapa))
	fmt.Printf("Soma mapa2: %.2f\n", Soma(mapa2))
	fmt.Printf("Soma mapa3: %d\n", Soma(mapa3))

	fmt.Println(Compara(10, 10))
}
