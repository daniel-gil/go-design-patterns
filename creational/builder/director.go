package builder

// ManufacturingDirector is the director component
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	// from the builder reproduces the build process to create a full vehicle object
	f.builder.SetSeats().SetStructure().SetWheels().SetColor().SetTopSpeed()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
