package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string, 3)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Println(fmt.Sprintf("sending hello %d", i))
			ch <- fmt.Sprintf("hello %d", i)
		}
		close(ch)
	}()

	//main goroutine waits here
	/*for i := 1; i <= 3; i++ {
		value := <-ch
		fmt.Println("received value:", value)
	}*/
	/*for {
		if msg, ok := <-ch; ok {
			fmt.Println("received value:", msg)
		} else {
			break
		}
	}*/
	for msg := range ch {
		fmt.Println("received value:", msg)
	}
	wg.Wait()
}
