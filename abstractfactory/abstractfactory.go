package abstractfactory

import (
	"errors"
	"fmt"
)

type VehicleAbstracktFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized", f))
	}
}
