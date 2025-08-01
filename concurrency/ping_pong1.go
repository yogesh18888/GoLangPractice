package main

import (
	"fmt"
	"sync"
)

// this program will have a deadlock! use buffer channel instead
func main() {
	ping := make(chan string)
	pong := make(chan string)
	wg := sync.WaitGroup{}
	exchanges := 5
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= exchanges; i++ {
			msg := <-ping
			fmt.Println("msg received:", msg)
			pong <- fmt.Sprintf("pong %d", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= exchanges; i++ {
			msg := <-pong
			fmt.Println("msg received:", msg)
			//at last iteration there will be deadlock, since ping goroutine's for loop is already over
			//and no receiver is ready for this last iteration operation
			ping <- fmt.Sprintf("ping %d", i)
		}
	}()

	ping <- "ping 1"
	wg.Wait()
	fmt.Println("Game over!")
}
