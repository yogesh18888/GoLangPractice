package main

import "fmt"

func describe(i interface{}) { // Accepts any type (empty interface)
	fmt.Printf("Value %v is of type ", i)
	switch v := i.(type) { // 'v' will hold the value of the concrete type
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Printf("unknown type %T\n", v) // %T prints the type
	}
}

func main() {
	fmt.Println("\n--- Type Switch ---")
	describe(42)
	describe("hello")
	describe(true)
	describe(3.14)
	describe([]int{1, 2, 3}) // A slice of ints
	describe(nil)            // The type of a nil interface is nil
}
