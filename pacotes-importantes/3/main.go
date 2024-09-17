package main

func main() {

	defer println("Primeira linha")
	println("Segunda linha")
	println("Terceira linha")

	// req, err := http.Get("https://www.google.com.br")
	// if err != nil {
	// 	panic(err)
	// }

	// defer func(Body io.ReadCloser) {
	// 	err := Body.Close()
	// 	if err != nil {
	// 		println(err)
	// 		return
	// 	}
	// }(req.Body)
	// res, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(res))

}
