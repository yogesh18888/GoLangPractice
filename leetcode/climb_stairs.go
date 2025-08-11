package main

import "fmt"

func main() {
	n := 6
	fmt.Printf("Number of ways to climb %d stairs: %d\n", n, climbStairs(n))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// Only store the last two values (space optimized)
	first, second := 1, 2

	for i := 3; i <= n; i++ {
		current := first + second
		first = second
		second = current
	}

	return second
}
