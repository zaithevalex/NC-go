package main

import (
	"fmt"
	"testing"
)

func IsNegativeNumber(number int) bool {
	return number < 0
}

func TestIsNegativeNumber(testing *testing.T) {
	answer := IsNegativeNumber(-1)
	if answer != true {
		testing.Errorf("IsNegativeNumber(-1) = %v, want %v", answer, true)
	}
}

func TestIsNegativeNumberTableDriven(t *testing.T) {
	tests := []struct {
		number int
		want   bool
	}{
		{1, false},
		{0, false},
		{-1, true},
	}
	for _, test := range tests {
		testName := fmt.Sprintf("%d", test.number)
		t.Run(testName, func(t *testing.T) {
			answer := IsNegativeNumber(test.number)
			if answer != test.want {
				t.Errorf("got %v, want %v", answer, test.want)
			}
		})
	}
}
