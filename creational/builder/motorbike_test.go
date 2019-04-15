package builder

import "testing"

func TestMotorbike(t *testing.T) {
	director := ManufacturingDirector{}

	motorbikeBuilder := &MotorbikeBuilder{}
	director.SetBuilder(motorbikeBuilder)
	director.Construct()

	motorbike := motorbikeBuilder.GetVehicle()

	const (
		defaultMotorbikeWeels     = 2
		defaultMotorbikeStructure = "Motorbike"
		defaultMotorbikeSeats     = 2
		defaultMotorbikeColor     = WHITE
		defaultMotorbikeTopSpeed  = 50
	)

	if motorbike.Wheels != defaultMotorbikeWeels {
		t.Errorf("expected %d wheels on a motorbike; got %d", defaultMotorbikeWeels, motorbike.Wheels)
	}
	if motorbike.Structure != defaultMotorbikeStructure {
		t.Errorf("expected '%s' structure on a motorbike; got %s", defaultMotorbikeStructure, motorbike.Structure)
	}
	if motorbike.Seats != defaultMotorbikeSeats {
		t.Errorf("expected %d seats on a motorbike; got %d", defaultMotorbikeSeats, motorbike.Seats)
	}
	if motorbike.Color != defaultMotorbikeColor {
		t.Errorf("expected %s color on a motorbike; got %s", defaultMotorbikeColor, motorbike.Color)
	}
	if motorbike.TopSpeed != defaultMotorbikeTopSpeed {
		t.Errorf("expected %d top speed on a motorbike; got %d", defaultMotorbikeTopSpeed, motorbike.TopSpeed)
	}
}
