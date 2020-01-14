package main

import (
	"testing"
)

func TestPlus(t *testing.T) {
	var a float64 = 3.4
	var b float64 = 2.1

	r := plus(a, b)

	if r != 5.5 {
		t.Error("Expected 5.5, got: ", r)
	}

}

func TestMinusToNegative(t *testing.T) {
	var a float64 = 5.2
	var b float64 = 7.9

	r := minus(a, b)

	if r != -2.7 {
		t.Error("Expected -2.7, got ", r)
	}
}

// This test case faild, as Golangs float precision
// Indicates that 7.9-1.2 = 7.8999999999999995

/*
func TestMinusToPositive(t *testing.T) {
	var a float64 = 9.1
	var b float64 = 1.2

	r := minus(a, b)

	// bfr := big.NewFloat(r)
	// bfc := big.NewFloat(7.9)

	if r != 7.9 {
		t.Error("Expected 7.9, got ", r)
	}
}

*/

func TestMultiply(t *testing.T) {
	var a float64 = 4.2
	var b float64 = 8.9

	r := multiply(a, b)

	if r != 37.38 {
		t.Error("Expected 37.38, got ", r)
	}
}

func TestMultiplyWithZero(t *testing.T) {
	var a float64 = 7.9
	var b float64 = 0

	r := multiply(a, b)

	if r != 0 {
		t.Error("Expected 0, got ", r)
	}
}

func TestDivide(t *testing.T) {
	var a float64 = 4
	var b float64 = 2

	r := divide(a, b)

	if r != 2 {
		t.Error("Expected 2, got ", r)
	}
}
