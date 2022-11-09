package mathematics

import "testing"

const erroPadrao = "the expected value %v, but the result found was: %v."

func TestAverage(t *testing.T) {
	expectedValue := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(erroPadrao, expectedValue, value)
	}
}

func TestAverage1(t *testing.T) {
	expectedValue1 := 8.28
	value := Average(7.28, 9.2, 7.2, 6.9)

	if value != expectedValue1 {
		t.Errorf(erroPadrao, expectedValue1, value)
	}
}

func TestAverage2(t *testing.T) {
	expectedValue2 := 9.00
	value := Average(9.28, 9.2, 9.8, 9.7)

	if value != expectedValue2 {
		t.Errorf(erroPadrao, expectedValue2, value)
	}
}
