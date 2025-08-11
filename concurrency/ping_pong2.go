package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Game Started!")
	pingCh := make(chan string)
	pongCh := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go ping(pingCh, pongCh, &wg)
	go pong(pingCh, pongCh, &wg)
	wg.Wait()
	fmt.Println("Game Over!")
}

func ping(pingCh chan<- string, pongChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		pingCh <- "ping"
		msg := <-pongChan
		fmt.Println("ping message received:", msg)
	}
}

func pong(pingCh <-chan string, pongCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		msg := <-pingCh
		fmt.Println("pong message received:", msg)
		pongCh <- "pong"
	}
}
