package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Olá mundo! teste"))
	if err != nil {
		panic(err)
	}
	fmt.Println(tamanho)
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)

	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// remover o arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
