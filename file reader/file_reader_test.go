package main

import (
	"os"
	"testing"
)

func TestFileReader(t *testing.T) {
	file := "./file.txt"
	_, err := os.ReadFile(file)

	if err != nil {
		t.Errorf("Failed, Got nil instead of %v", file)
		os.Exit(1)
	}

}
