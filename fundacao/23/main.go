package main

func main() {
	// for i := 0; i < 10; i++ {
	// 	print(i)
	// }

	numeros := []string{"um", "dois", "tres"}

	for index, value := range numeros {
		println(index, value)
	}

	//ignorando index
	for _, value := range numeros {
		println(value)
	}
	//ignorando valor
	for index, _ := range numeros {
		println(index)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Hello, World!!")
	}

}
