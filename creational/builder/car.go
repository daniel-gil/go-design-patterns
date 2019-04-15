package builder

// CarBuilder is a concrete builder component
// Implements the builder interfaces for the type car
type CarBuilder struct {
	vehicleProduct VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.vehicleProduct.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.vehicleProduct.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.vehicleProduct.Structure = "Car"
	return c
}

func (c *CarBuilder) SetColor() BuildProcess {
	c.vehicleProduct.Color = WHITE
	return c
}

func (c *CarBuilder) SetTopSpeed() BuildProcess {
	c.vehicleProduct.TopSpeed = 120
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.vehicleProduct
}
