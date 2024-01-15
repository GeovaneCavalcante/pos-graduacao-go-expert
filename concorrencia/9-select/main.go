package main

import (
	"fmt"
	"time"
)

type Message struct {
	ID  int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	go func() {
		time.Sleep(2 * time.Second)
		msg := Message{1, "hello rabbitMQ"}
		c1 <- msg

	}()

	go func() {
		time.Sleep(1 * time.Second)
		msg := Message{1, "hello kafka"}
		c2 <- msg

	}()

	// for i := 0; i < 2; i++ {
	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("received from RebbitMQ: %s\n", msg1.Msg)
		case msg2 := <-c2:
			fmt.Printf("received from kafka: %s\n", msg2.Msg)

		case <-time.After(3 * time.Second):
			println("timeout")
			// default:
			// 	println("default")
		}
	}

}
