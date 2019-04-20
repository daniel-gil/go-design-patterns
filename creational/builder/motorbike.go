package builder

// MotorbikeBuilder is a concrete builder component
// Implements the builder interfaces for the type motorbike
type MotorbikeBuilder struct {
	vehicleProduct VehicleProduct
}

func (m *MotorbikeBuilder) SetWheels(wheels int) BuildProcess {
	m.vehicleProduct.Wheels = wheels
	return m
}

func (m *MotorbikeBuilder) SetDefaultWheels() BuildProcess {
	return m.SetWheels(2)
}

func (m *MotorbikeBuilder) SetSeats(seats int) BuildProcess {
	m.vehicleProduct.Seats = seats
	return m
}

func (m *MotorbikeBuilder) SetDefaultSeats() BuildProcess {
	return m.SetSeats(2)
}

func (m *MotorbikeBuilder) SetStructure(structure string) BuildProcess {
	m.vehicleProduct.Structure = structure
	return m
}

func (m *MotorbikeBuilder) SetDefaultStructure() BuildProcess {
	return m.SetStructure("Motorbike")
}

func (m *MotorbikeBuilder) SetColor(color Color) BuildProcess {
	m.vehicleProduct.Color = color
	return m
}

func (m *MotorbikeBuilder) SetDefaultColor() BuildProcess {
	return m.SetColor(WHITE)
}

func (m *MotorbikeBuilder) SetTopSpeed(speed int) BuildProcess {
	m.vehicleProduct.TopSpeed = speed
	return m
}

func (m *MotorbikeBuilder) SetDefaultTopSpeed() BuildProcess {
	return m.SetTopSpeed(50)
}

func (m *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return m.vehicleProduct
}
