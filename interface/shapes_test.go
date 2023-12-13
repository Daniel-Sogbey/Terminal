package main

import "testing"

func TestGetArea(t *testing.T) {
	s := square{
		sideLength: 5,
	}

	if s.getArea() < 0 {
		t.Errorf("Expected area to be 10, but got %v", s.getArea())
	}
}
