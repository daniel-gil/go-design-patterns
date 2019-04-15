package builder

// BikeBuilder is a concrete builder component
// Implements the builder interfaces for the type bike
type BikeBuilder struct {
	vehicleProduct VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.vehicleProduct.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.vehicleProduct.Seats = 1
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.vehicleProduct.Structure = "Bike"
	return b
}

func (b *BikeBuilder) SetColor() BuildProcess {
	b.vehicleProduct.Color = WHITE
	return b
}

func (b *BikeBuilder) SetTopSpeed() BuildProcess {
	b.vehicleProduct.TopSpeed = 20
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.vehicleProduct
}
