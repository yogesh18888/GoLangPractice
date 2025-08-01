package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	msgChan := make(chan string)

	wg.Add(1)
	go receive(msgChan, &wg)

	wg.Add(1)
	go send(msgChan, &wg)

	wg.Wait()
}

func send(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "hello"
}

func receive(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	fmt.Println("received message:", msg)
}
