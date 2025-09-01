package main

import (
	"context"
	"fmt"
	"time"
)

// A worker function that simulates a long-running task.
// It uses a context to know when to stop.
func worker(ctx context.Context, name string, duration time.Duration) {
	fmt.Printf("Worker %s started, running for %v...\n", name, duration)

	select {
	case <-time.After(duration):
		// This case is for the worker's natural completion.
		// If the context is not canceled, the task will finish.
		fmt.Printf("Worker %s finished its task normally.\n", name)
	case <-ctx.Done():
		// This case is triggered when the context is canceled (by timeout or manually).
		// The worker stops its task and cleans up.
		fmt.Printf("Worker %s received cancellation signal. Reason: %v\n", name, ctx.Err())
	}
}

func main() {
	fmt.Println("Starting main program.")

	// Create a context with a timeout of 2 seconds.
	// The cancel function is used to release resources associated with the context.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// It's crucial to call cancel() when you're done with the context
	// to prevent a goroutine leak.
	defer cancel()

	// --- Scenario 1: Timeout Occurs ---
	fmt.Println("\n--- Scenario 1: Worker takes longer than timeout ---")
	// Start a worker that needs 3 seconds to complete.
	go worker(ctx, "Long-Running Worker", 3*time.Second)

	// We'll wait here for a bit to see the outcome.
	time.Sleep(2500 * time.Millisecond) // Wait 2.5 seconds to see the cancellation
	fmt.Println("Main program continues after waiting.")

	// The `defer cancel()` will eventually be called, but the timeout
	// already triggered the cancellation.

	// --- Scenario 2: Worker Finishes within Timeout ---
	fmt.Println("\n--- Scenario 2: Worker finishes before timeout ---")
	// Create a new context for this scenario.
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()

	// Start a worker that needs only 1 second to complete.
	go worker(ctx2, "Short-Running Worker", 1*time.Second)

	// Wait for the worker to finish.
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Main program continues after waiting.")
}
