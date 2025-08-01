package main

import (
	"fmt"
	"time"
)

func main() {
	/*myChannel := make(chan int)
	go produce(myChannel)

	for v := range myChannel {
		fmt.Println("channel value: ", v)
	}

	for {
		v, ok := <-myChannel
		if !ok {
			fmt.Println("channel closed")
			break
		}
		fmt.Println("channel value: ", v)
	}*/
	ch := make(chan int)
	go func() {
		chValue := <-ch
		fmt.Println("value:", chValue)
	}()
	ch <- 1
	fmt.Println("hello")
}

func produce(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}
