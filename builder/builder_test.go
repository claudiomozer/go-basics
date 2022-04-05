package builder

import "testing"

func TestBuilderPattern(t *testing.T) {

	manufactoringDirector := ManufactoringDirector{}

	carBuilder := &CarBuilder{}
	manufactoringDirector.SetBuilder(carBuilder)
	manufactoringDirector.Construct()

	car := carBuilder.GetVehicle()

	if car.Seats != 5 {
		t.Errorf("A car must have 5 seats, and it has %d\n", car.Seats)
	}

	if car.Wheels != 4 {
		t.Errorf("A car must have 4 wheels, and it has %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("A car structure must be 'Car', and it has %s\n", car.Structure)
	}

	bikeBuilder := &BikeBuilder{}
	manufactoringDirector.SetBuilder(bikeBuilder)
	manufactoringDirector.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Seats != 2 {
		t.Errorf("A bike must have 2 seats, and it has %d\n", bike.Seats)
	}

	if bike.Wheels != 2 {
		t.Errorf("A bike must have 2 wheels, and it has %d\n", bike.Wheels)
	}

	if bike.Structure != "Motorbike" {
		t.Errorf("A bike structure must be 'Motorbike', and it has %s\n", bike.Structure)
	}
}
