package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 2)
}

func worker(ctx context.Context) {
	fmt.Println("started goroutine")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker done:", ctx.Err())
			return
		default:
			fmt.Println("working...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
