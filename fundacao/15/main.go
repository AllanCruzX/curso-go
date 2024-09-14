package main

func main() {

	a := 10
	var ponteiro *int = &a
	b := &a

	println(&a)
	println(ponteiro)
	println(b)

}
