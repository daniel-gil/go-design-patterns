package builder

import "testing"

func TestCar(t *testing.T) {
	director := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	director.SetBuilder(carBuilder)
	director.Construct()

	car := carBuilder.GetVehicle()

	const (
		defaultCarWeels     = 4
		defaultCarStructure = "Car"
		defaultCarSeats     = 5
		defaultCarColor     = WHITE
		defaultCarTopSpeed  = 120
	)

	if car.Wheels != defaultCarWeels {
		t.Errorf("expected %d wheels on a car; got %d", defaultCarWeels, car.Wheels)
	}
	if car.Structure != defaultCarStructure {
		t.Errorf("expected '%s' structure on a car; got %s", defaultCarStructure, car.Structure)
	}
	if car.Seats != defaultCarSeats {
		t.Errorf("expected %d seats on a car; got %d", defaultCarSeats, car.Seats)
	}
	if car.Color != defaultCarColor {
		t.Errorf("expected %s color on a car; got %s", defaultCarColor, car.Color)
	}
	if car.TopSpeed != defaultCarTopSpeed {
		t.Errorf("expected %d top speed on a car; got %d", defaultCarTopSpeed, car.TopSpeed)
	}
}
