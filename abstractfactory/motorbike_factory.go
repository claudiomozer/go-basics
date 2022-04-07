package abstractfactory

import (
	"errors"
	"fmt"
)

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFacktory struct{}

func (m *MotorbikeFacktory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Motorbike of type %d not recognized\n", v))

	}
}
