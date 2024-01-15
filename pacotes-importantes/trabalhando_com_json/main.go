package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 123, Saldo: 1000}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	jsonPuro := `{"n":123,"s":1000}`

	var contaX Conta
	err = json.Unmarshal([]byte(jsonPuro), &contaX)
	if err != nil {
		panic(err)
	}

	fmt.Println(contaX)

}
