package main

import (
	"fmt"
	"time"
)

// let's play ping pong for 5 seconds
func main() {
	pingCh := make(chan string)
	pongCh := make(chan string)

	go func() {
		for {
			msg := <-pingCh
			fmt.Println("ping received:", msg)
			time.Sleep(500 * time.Millisecond)
			pongCh <- "pong"
		}
	}()

	go func() {
		for {
			msg := <-pongCh
			fmt.Println("pong received:", msg)
			time.Sleep(500 * time.Millisecond)
			pingCh <- "ping"
		}
	}()

	// let's start
	pingCh <- "ping"

	time.Sleep(time.Second * 3)
	fmt.Println("Game Over!")
	//once main goroutine finishes, the entire go program exits,
	//and all other goroutines are forcefully terminated
}
