package main

import "fmt"

func main() {
	p1 := Person{
		Name: "John Doe",
		Age:  20,
	}
	fmt.Println(p1)
	fmt.Println(&p1)

	ptrP1 := &p1
	fmt.Println((*ptrP1).Name, ":", ptrP1.Age)
	ptrP1.Age = 30
	fmt.Println((*ptrP1).Name, ":", ptrP1.Age)
}

type Person struct {
	Name string
	Age  int
}
