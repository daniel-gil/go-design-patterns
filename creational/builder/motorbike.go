package builder

// MotorbikeBuilder is a concrete builder component
// Implements the builder interfaces for the type motorbike
type MotorbikeBuilder struct {
	vehicleProduct VehicleProduct
}

func (m *MotorbikeBuilder) SetWheels() BuildProcess {
	m.vehicleProduct.Wheels = 2
	return m
}

func (m *MotorbikeBuilder) SetSeats() BuildProcess {
	m.vehicleProduct.Seats = 2
	return m
}

func (m *MotorbikeBuilder) SetStructure() BuildProcess {
	m.vehicleProduct.Structure = "Motorbike"
	return m
}

func (m *MotorbikeBuilder) SetColor() BuildProcess {
	m.vehicleProduct.Color = WHITE
	return m
}

func (m *MotorbikeBuilder) SetTopSpeed() BuildProcess {
	m.vehicleProduct.TopSpeed = 50
	return m
}

func (m *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return m.vehicleProduct
}
