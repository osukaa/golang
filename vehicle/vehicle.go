package vehicle

// Vehicle is nice
type Vehicle interface {
	Move()
}

// Move any vehicle
func Move(vehicle Vehicle) {
	vehicle.Move()
}
