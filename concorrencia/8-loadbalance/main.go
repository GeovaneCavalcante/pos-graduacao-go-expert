package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for i := range data {
		fmt.Printf("Worker %d received %d\n", workerID, i)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtdWorker := 99

	for i := 1; i < qtdWorker; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}
