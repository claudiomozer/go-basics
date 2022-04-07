package abstractfactory

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}
