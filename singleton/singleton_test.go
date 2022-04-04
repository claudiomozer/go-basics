package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("expected ponter to Singleton")
	}
	expectedCounter := counter1
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 bu it's %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expect the same instace in counter to but it got a different one")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling AddOne() with the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}
