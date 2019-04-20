package builder

// BuildProcess is the abstract builder component
type BuildProcess interface {
	SetDefaultWheels() BuildProcess
	SetDefaultSeats() BuildProcess
	SetDefaultStructure() BuildProcess
	SetDefaultColor() BuildProcess
	SetDefaultTopSpeed() BuildProcess

	SetWheels(wheels int) BuildProcess
	SetSeats(seats int) BuildProcess
	SetStructure(structure string) BuildProcess
	SetColor(color Color) BuildProcess
	SetTopSpeed(speed int) BuildProcess

	GetVehicle() VehicleProduct
}
