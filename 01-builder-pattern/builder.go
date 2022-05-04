package builder

type Wheel int

type Vehicle struct {
	Wheels    Wheel
	Seats     int
	Structure string
}

// Builder interface
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	GetVehicle() Vehicle
}

// Car struct
type Car struct {
	vehicle Vehicle
}

// SetWheels implement Builder
func (car *Car) SetWheels() Builder {
	car.vehicle.Wheels = 4
	return car
}

func (car *Car) SetSeats() Builder {
	car.vehicle.Seats = 4
	return car
}

func (car *Car) SetStructure() Builder {
	car.vehicle.Structure = "Car"
	return car
}

func (car *Car) GetVehicle() Vehicle {
	return car.vehicle
}

func (car *Car) Construct() {
	car.SetWheels().SetSeats().SetStructure() // 链式调用
}
