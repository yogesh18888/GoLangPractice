package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			mutex.Lock()
			counter = counter + 2
			mutex.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("counter:", counter)
}
