package builder

// CarBuilder is a concrete builder component
// Implements the builder interfaces for the type car
type CarBuilder struct {
	vehicleProduct VehicleProduct
}

func (c *CarBuilder) SetWheels(wheels int) BuildProcess {
	c.vehicleProduct.Wheels = wheels
	return c
}

func (c *CarBuilder) SetDefaultWheels() BuildProcess {
	return c.SetWheels(4)
}

func (c *CarBuilder) SetSeats(seats int) BuildProcess {
	c.vehicleProduct.Seats = seats
	return c
}

func (c *CarBuilder) SetDefaultSeats() BuildProcess {
	return c.SetSeats(5)
}

func (c *CarBuilder) SetStructure(structure string) BuildProcess {
	c.vehicleProduct.Structure = structure
	return c
}

func (c *CarBuilder) SetDefaultStructure() BuildProcess {
	return c.SetStructure("Car")
}

func (c *CarBuilder) SetColor(color Color) BuildProcess {
	c.vehicleProduct.Color = color
	return c
}

func (c *CarBuilder) SetDefaultColor() BuildProcess {
	return c.SetColor(WHITE)
}

func (c *CarBuilder) SetTopSpeed(speed int) BuildProcess {
	c.vehicleProduct.TopSpeed = speed
	return c
}

func (c *CarBuilder) SetDefaultTopSpeed() BuildProcess {
	return c.SetTopSpeed(120)
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.vehicleProduct
}
