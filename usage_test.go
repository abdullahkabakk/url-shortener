package main

import (
	"testing"
)

func TestBToMB(t *testing.T) {
	b := uint64(1024 * 1024)

	result := bToMb(b)

	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestPrintMemoryUsage(t *testing.T) {
	printMemoryUsage()

	// No need to test the output of this function
	// It's just for debugging purposes

}
