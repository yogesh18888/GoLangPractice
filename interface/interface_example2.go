package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{
		radius: 5,
	}
	fmt.Println(circle.Area())
	circle.Draw()

	rectangle := Rectangle{
		width:  5,
		height: 5,
		length: 5,
	}
	fmt.Println(rectangle.Area())
	rectangle.Draw()

	DetectShape(circle)
	DetectShape(rectangle)
}

type Shape interface {
	Area() float64
	Draw()
}

type Rectangle struct {
	length, width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width * r.height
}
func (r Rectangle) Draw() {
	fmt.Println("Rectangle drawn")
}

func (c Circle) Area() float64 {
	return c.radius * math.Pi
}
func (c Circle) Draw() {
	fmt.Println("Circle drawn")
}

func DetectShape(shape Shape) {
	if s, ok := shape.(Circle); ok {
		fmt.Println("given shape is Circle with Radius: ", s.radius)
	} else if s, ok := shape.(Rectangle); ok {
		fmt.Println("Given shape is Rectangle with Length: ", s.length)
	}
}
