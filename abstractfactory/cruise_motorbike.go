package abstractfactory

type CruiseMotorbike struct{}

func (s *CruiseMotorbike) GetWheels() int {
	return 2
}

func (s *CruiseMotorbike) GetSeats() int {
	return 2
}

func (s *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}
