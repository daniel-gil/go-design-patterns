package builder

// BuildProcess is the abstract builder component
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	SetColor() BuildProcess
	SetTopSpeed() BuildProcess
	GetVehicle() VehicleProduct
}
