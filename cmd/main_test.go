package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sum(2, 3)
	expected := 5
	
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}


func TestSubtract(t *testing.T) {
	result := subtract(2, 3)
	expected := -1

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
    result := multiply(2, 3)
    expected := 6
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}

func TestDivide(t *testing.T) {
    result, err := divide(6, 3)
    expected := 2
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    } else if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
    
    _, err = divide(6, 0)
    if err == nil {
        t.Error("Expected error for division by zero, got nil")
    }
}
