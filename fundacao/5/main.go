package main

import (
	"fmt"
)

const a = "Hello, word!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Allan"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	//fmt.Println(len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("O valor do inidice é %d e o valor é %d \n", i, v)
	}
}
