package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("main started")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userId", "yogesh123")
	greetUser(ctx)
}

func greetUser(ctx context.Context) {
	name := ctx.Value("userId")
	if name != nil {
		fmt.Println("Hello ", name)
	} else {
		fmt.Println("Hello guest")
	}
}
