package main

import (
	"fmt"
	"sync"
)

// print input in same order using goroutines
func main() {
	input := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
	result := make([]string, len(input))
	wg := sync.WaitGroup{}
	for i, v := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[i] = v
		}()
	}
	wg.Wait()
	fmt.Println(result)

	// using two goroutines now
	fmt.Println("printing using two goroutines with unbuffered channels")
	result1 := make([]string, len(input))
	length := len(input)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < length; i += 2 {
			<-ch1 // wait for signal
			result1[i] = input[i]
			ch2 <- true // signal goroutine 2
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < length; i += 2 {
			<-ch2
			result1[i] = input[i]
			if i+1 < len(input) {
				fmt.Println("entered here:", result1[i])
				ch1 <- true // signal goroutine 1
			}
		}
	}()
	ch1 <- true
	wg.Wait()
	fmt.Println(result1)
}
