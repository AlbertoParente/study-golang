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
