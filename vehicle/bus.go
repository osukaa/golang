package vehicle

import "fmt"

// Bus is a type of vehicle
type Bus struct {
}

// Move the bus
func (b *Bus) Move() {
	fmt.Print("Moving moving moving")
}

// Forward moves the bus forward
func (b *Bus) Forward() {
	fmt.Print("I am moving forward")
}

// Backward moves the bus backwards
func (b *Bus) Backward() {
	fmt.Print("I am moving backward")
}
