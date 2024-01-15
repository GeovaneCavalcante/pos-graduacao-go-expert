package main

import "fmt"

const a = "Hello world"

type ID int

var b bool = true
var c int = 10
var d string = "Geovane"
var e float64 = 1.2

var f ID = 1

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é o valor %d\n", i, v)
	}
}
