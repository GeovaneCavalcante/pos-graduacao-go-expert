package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		println(url)
		req, err := http.Get(url)
		if err != nil {
			fmt.Println(os.Stderr, "error")
		}
		defer req.Body.Close()

		res, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("erro ao ler o arquivo")
		}

		var endereco Endereco
		err = json.Unmarshal(res, &endereco)
		if err != nil {
			fmt.Println("erro ao converter para struct")
		}

		fmt.Printf("%+v\n", endereco)

		file, err := os.Create("endereco.text")
		if err != nil {
			fmt.Println("erro ao criar arquivo")
		}

		defer file.Close()
		file.WriteString(fmt.Sprintf("%+v\n", endereco))

	}
}
