package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)
	// Thread 2
	go task("a", &waitGroup)
	// Thread 3
	go task("b", &waitGroup)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "c")
			time.Sleep(1 * time.Second)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
