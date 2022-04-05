package builder

type BuildVehicle interface {
	SetWheels() BuildVehicle
	SetSeats() BuildVehicle
	SetStructure() BuildVehicle
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufactoringDirector struct {
	builder BuildVehicle
}

func (f *ManufactoringDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufactoringDirector) SetBuilder(builder BuildVehicle) {
	f.builder = builder
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildVehicle {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildVehicle {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildVehicle {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildVehicle {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildVehicle {
	c.v.Seats = 2
	return c
}

func (c *BikeBuilder) SetStructure() BuildVehicle {
	c.v.Structure = "Motorbike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}
