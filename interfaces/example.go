package main

import (
	"fmt"
	"math"
)

func main() {
	var greeter Greeter = EnglishGreeter{
		Name: "Hello",
	}
	greeter.sayHello()

	greeter = MarathiGreeter{
		Name: "Namaste",
	}
	greeter.sayHello()
	fmt.Println(cap([]int{1, 2, 3}))
	var p Person
	p.name = "yogesh"
	p.age = 37
	p.gender = "male"
	fmt.Println(p)

	p2 := new(Person)
	p2.name = "yogesh2"
	fmt.Println(p2)

	p3 := Person{name: "yogesh3", age: 37, gender: "male"}
	fmt.Println(p3)

	var s Shape = &Circle{radius: 5}
	fmt.Println(s.area())
	if c, ok := s.(*Circle); ok {
		fmt.Println("circle result:", c.result)
	}

	s = Rectangle{length: 5, height: 5}
	fmt.Println(s.area())
}

type Person struct {
	name   string
	age    int
	gender string
}

type Greeter interface {
	sayHello()
}

type EnglishGreeter struct {
	Name string
}

type MarathiGreeter struct {
	Name string
}

func (eg EnglishGreeter) sayHello() {
	fmt.Println(eg.Name, "Yogesh!")
}

func (mg MarathiGreeter) sayHello() {
	fmt.Println(mg.Name, "Yogesh!")
}

// shape example
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
	result float64
}
type Rectangle struct {
	length float64
	height float64
}

func (c *Circle) area() float64 {
	res := math.Pi * c.radius * c.radius
	c.result = res
	return res
}
func (r Rectangle) area() float64 {
	return r.length * r.height
}
