package main

import (
	"testing"
)

func TestGuessIt(t *testing.T) {
	tests := []struct {
		input    []float64
		expectedLower int
		expectedUpper int
	}{
		{[]float64{1.5, 2.5, 3.5, 4.5}, 1, 5}, 
		{[]float64{10, 20, 30, 40, 50}, 2, 58},
	}

	for _, test := range tests {
		lower, upper := guessit(test.input)
		if lower != test.expectedLower || upper != test.expectedUpper {
			t.Errorf("For input %v, expected (%d, %d) but got (%d, %d)", test.input, test.expectedLower, test.expectedUpper, lower, upper)
		}
	}
}
