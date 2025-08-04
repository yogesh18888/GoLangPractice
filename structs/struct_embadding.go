package main

import "fmt"

// Base struct representing a Vehicle
type Vehicle struct {
	Speed int
}

// Method for Vehicle
func (v *Vehicle) accelerate() {
	v.Speed += 10
	fmt.Println(fmt.Sprintf("Vehicle accelerated. Speed: %d\n", v.Speed))
}

// Embedding the Vehicle struct in the Car struct
type Car struct {
	Vehicle // Embedded struct (composition)
	Make    string
	Model   string
}

func main() {
	myCar := Car{
		Make:  "Toyota",
		Model: "Camry",
		Vehicle: Vehicle{
			Speed: 0,
		},
	}
	// We can directly access fields and methods of the embedded Vehicle
	// through the Car struct instance (e.g., myCar.accelerate())
	fmt.Println("Initial speed:", myCar.Speed)
	myCar.accelerate()
	myCar.accelerate()
}
