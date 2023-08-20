package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}

func TestDivisionTestify(t *testing.T) {
	x := 6
	y := 3

	var expected float32 = 2

	actual, err := divide(float32(x), float32(y))

	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, expected, actual)
}

func TestDivision(t *testing.T) {
	testCases := []test{
		{"valid-data", 100.0, 10.0, 10.0, false},
		{"invalid-data", 100.0, 0.0, 0, true},
	}
	for _, tt := range testCases {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}
