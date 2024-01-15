package main

import "fmt"

func recebe(nome string, hello chan<- string) {
	hello <- nome
}
func ler(hello <-chan string) {
	fmt.Println(<-hello)
}

func main() {
	hello := make(chan string)
	go recebe("hello", hello)
	ler(hello)
}
