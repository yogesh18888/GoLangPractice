package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	queue []int
	mu    sync.Mutex
	cond  = sync.NewCond(&mu) // Associate with a Mutex
)

func producer(id int) {
	for i := 0; i < 5; i++ {
		mu.Lock()
		queue = append(queue, i)
		fmt.Printf("Producer %d added: %d, Queue: %v\n", id, i, queue)
		cond.Signal() // Signal one waiting consumer
		mu.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(id int) {
	for {
		mu.Lock()
		for len(queue) == 0 { // Loop to handle spurious wakeups
			fmt.Printf("Consumer %d: Queue empty, waiting...\n", id)
			cond.Wait() // Release lock, wait, re-acquire lock
		}
		item := queue[0]
		queue = queue[1:]
		fmt.Printf("Consumer %d consumed: %d, Queue: %v\n", id, item, queue)
		mu.Unlock()
		time.Sleep(200 * time.Millisecond) // Simulate consumption time
	}
}

func main() {
	go producer(1)
	//go producer(2)
	go consumer(1)
	//go consumer(2)

	// Let the program run for a while, then exit.
	// In a real app, you'd have proper shutdown mechanisms.
	time.Sleep(5 * time.Second)
	fmt.Println("Main goroutine exiting.")
}
