package main

import "fmt"

func main() {
	fmt.Println(sumAll(10, 20, 30))
	nums := []int{10, 20, 30, 40}
	fmt.Println(sumAll(nums...))

	//anonymous function
	add := func(a, b int) int { return a + b }
	fmt.Println(add(10, 10))

	// immediately invoked anonymous function
	func() {
		fmt.Println("this is immediately invoked anonymous function")
	}()

	//pointers as func arguments
	num1 := 1
	fmt.Println("before:", num1)
	incrementByValue(num1)
	fmt.Println("after:", num1)

	num2 := 2
	fmt.Println("before:", num2)
	incrementByPtr(&num2)
	fmt.Println("after:", num2)
}

func incrementByValue(x int) {
	x++
	fmt.Println("incrementByValue:", x)
}

func incrementByPtr(x *int) {
	*x++
	fmt.Println("incrementByPtr:", *x)
}

// variadic functions
func sumAll(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
