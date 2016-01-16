package main

import (
	"hello-world/util"
	"hello-world/vehicle"
)

func main() {
	util.Helloworld()
	bus := &vehicle.Schoolbus{}
	vehicle.Move(bus)
}
