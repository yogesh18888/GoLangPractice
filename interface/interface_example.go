package main

import "fmt"

type Greeter interface {
	SayHello() string
}

func main() {
	var g Greeter = EnglishSpeaker{Name: "John"}
	fmt.Println(g.SayHello())

	var g1 Greeter = MarathiSpeaker{Name: "Yogesh"}
	fmt.Println(g1.SayHello())

	// Type assertion: Check if g1 is actually a MarathiSpeaker
	if ms, ok := g1.(MarathiSpeaker); ok {
		ms.Test()
	}
}

type EnglishSpeaker struct {
	Name string
}

func (es EnglishSpeaker) SayHello() string {
	return "Hello! " + es.Name
}

type MarathiSpeaker struct {
	Name string
}

func (m MarathiSpeaker) SayHello() string {
	return "Namaste! " + m.Name
}

// Test This method is specific to MarathiSpeaker struct. This is not a common behaviour
// so in order to call this method we need to assert it
func (m MarathiSpeaker) Test() {
	fmt.Println("Test")
}
