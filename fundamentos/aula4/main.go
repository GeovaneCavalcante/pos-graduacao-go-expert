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
	fmt.Printf("o tipo de E é %T\n", f)
}
