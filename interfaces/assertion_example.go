package main

import "fmt"

func main() {
	var str *string
	myStr := "Hello World!"
	str = &myStr
	var i interface{} = str
	s, ok := i.(*string)
	fmt.Println(*s, ok)

	s = i.(*string)
	fmt.Println(*s)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64)
	var x int = 10
	var ptr *int = &x // ptr now holds the memory address of x
	fmt.Println(*ptr) // Dereference ptr to get the value, output: 10

	var op func(int, int) int = add // op is now a variable of function type
	fmt.Println(op(2, 2))
}

func add(a, b int) int {
	return a + b
}
