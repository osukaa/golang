package vehicle

import "fmt"

// Schoolbus a type of bus
type Schoolbus struct {
	*Bus
}

// Move a schoolbus
func (b *Schoolbus) Move() {
	fmt.Print("I am a school bus")
}
