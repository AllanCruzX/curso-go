package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ddd         string `json:"ddd"`
}

func main() {
	for _, cep := range os.Args[1:] {
		resposta := getCep(cep)
		if resposta != nil {
			criaArquivo(*resposta)
			fmt.Printf("%#v", resposta)
		}
	}
}

func getCep(cep string) *Endereco {

	urlViaCep := "https://viacep.com.br/ws/" + cep + "/json/"

	requisicao, err := http.Get(urlViaCep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(requisicao.Body)
	resposta, err := io.ReadAll(requisicao.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler retorno: %v\n", err)
	}

	var endereco Endereco

	err = json.Unmarshal(resposta, &endereco)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}

	return &endereco

}

func criaArquivo(endereco Endereco) {

	file, err := os.Create("cidade.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, " Erro ao criar arquivo: %v\n", err)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%#v", endereco))
	if err != nil {
		fmt.Fprintf(os.Stderr, " Erro gravar arquivo: %v\n", err)
	}
}
