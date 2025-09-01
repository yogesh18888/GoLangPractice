package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context.Background() is your root context.
	// You create a cancelable child context with context.WithCancel().
	// Multiple goroutines share the same ctx and monitor <-ctx.Done() to know when to stop.
	// When you call cancel():
	//	The cancel signal propagates to all children of ctx.
	//	All waiting goroutines are notified simultaneously.
	//	ctx.Err() becomes context.Canceled
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("started goroutine1")
		<-ctx.Done()
		fmt.Println("context canceled, stopping goroutine1")
		//Think of <-ctx.Done() as:
		//"Wait at this line until someone rings a bell. Once the bell is rung (cancel is called), continue with the rest of the code."
		//Until that bell is rung, the goroutine waits patiently, doing nothing.
	}()

	go func() {
		fmt.Println("started goroutine2")
		<-ctx.Done()
		fmt.Println("context canceled, stopping goroutine2")
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("signalling all goroutines who are using ctx after 3 seconds")
	cancel() //All goroutines receive signal when this gets executed after 5 seconds from: <-ctx.Done()
	time.Sleep(time.Millisecond * 500)
}
