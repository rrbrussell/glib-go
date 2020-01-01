package glib

import (
	"math"
	"testing"
)

func TestUint64CheckedAdd(t *testing.T) {
	type testValue struct {
		a        uint64
		b        uint64
		expected bool
	}
	testValues := []testValue{testValue{math.MaxUint32, 1, false}, testValue{math.MaxUint64, math.MaxUint8, true}}
	for i := 0; i < len(testValues); i++ {
		_, result := Uint64CheckedAdd(testValues[i].a, testValues[i].b) // I am assuming that math.big has already been tested for correctness
		if result != testValues[i].expected {
			t.Errorf("I expected %t when testing Uint64CheckedAdd(%d, %d). I got %t", testValues[i].expected, testValues[i].a, testValues[i].b, result)
		}
	}
}

func TestUint64CheckedMul(t *testing.T) {
	type testValue struct {
		a        uint64
		b        uint64
		expected bool
	}
	testValues := []testValue{testValue{math.MaxUint32, 1, false}, testValue{math.MaxUint64, math.MaxUint64, true}}
	for i := 0; i < len(testValues); i++ {
		_, result := Uint64CheckedMul(testValues[i].a, testValues[i].b) // I am assuming that math.big has already been tested for correctness
		if result != testValues[i].expected {
			t.Errorf("I expected %t when testing Uint64CheckedMul(%d, %d). I got %t", testValues[i].expected, testValues[i].a, testValues[i].b, result)
		}
	}
}
