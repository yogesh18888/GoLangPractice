package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(1)
	go receive(&wg, ch)

	wg.Add(1)
	go send(&wg, ch)

	wg.Wait()

	fmt.Println("\nDifferences with escape sequences:")
	fmt.Println("Interpreted: \"Hello\\nWorld\" ->", "Hello\nWorld")
	fmt.Println("Raw:         `Hello\\nWorld` ->", `Hello\nWorld`)

	s := "hello"
	// s[0] = 'H' // This would be a compile-time error: cannot assign to s[0]
	s = s + " world" // This creates a *new* string and assigns it to 's'
	fmt.Println(s)

}

func send(wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	ch <- "Hello World"
	close(ch)
}

func receive(wg *sync.WaitGroup, ch <-chan string) {
	defer wg.Done()
	msg := <-ch
	fmt.Println(msg)
}
