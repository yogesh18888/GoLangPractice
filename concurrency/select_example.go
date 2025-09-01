package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "hello"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		ch2 <- "world"
	}()

	for i := 1; i <= 2; i++ {
		fmt.Println("iterating over ", i)
		select {
		case msg1 := <-ch1:
			fmt.Println("received value from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("received value from ch2:", msg2)
		case <-time.After(time.Second * 4):
			fmt.Println("no value received")
		}
	}
}
