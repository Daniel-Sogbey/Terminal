package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	mh := myHandler{}

	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Errorf("Type is not http.Handler. Got %t", v)
	}
}

func TestSessionLoad(t *testing.T) {
	mh := myHandler{}

	h := SessionLoad(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Errorf("Type is not http.Handler. Got %t", v)
	}
}
