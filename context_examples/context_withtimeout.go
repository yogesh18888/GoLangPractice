package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("started main")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel() // Clean up context even if not used
	go slowOperation(ctx)
	doWork(ctx)
	time.Sleep(time.Millisecond * 5000)
}

func slowOperation(ctx context.Context) {
	fmt.Println("started slowOperation")
	select {
	case <-time.After(time.Second * 4):
		fmt.Println("slow operation")
	case <-ctx.Done():
		fmt.Println("context expired:", ctx.Err())
		fmt.Println("cleaning the resources")
		// The ctx.Done() channel my not be triggered every time if operation finishes before deadline
	}
}

func doWork(ctx context.Context) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		if ctx.Err() != nil {
			fmt.Println("worker done:", ctx.Err())
			return
		}
		fmt.Println("working on:", i)
	}
	fmt.Println("worker done!")
}
