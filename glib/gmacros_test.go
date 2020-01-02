package glib

import "testing"

func TestClamp(t *testing.T) {
	testX := [][]float64{{1, -2, 4, 1}, {9000, 10000, 12000, 10000}, {-64, -76, -70, -70}}
	for i := 0; i < len(testX); i++ {
		result := Clamp(testX[i][0], testX[i][1], testX[i][2])
		if result != testX[i][3] {
			t.Errorf("I expected %f when testing Clamp(%f, %f, %f). I got %f.", testX[i][3], testX[i][0], testX[i][1], testX[i][2], result)
		}
	}
}

func TestApproxValue(t *testing.T) {
	type testValue struct {
		a        float64
		b        float64
		epsilon  float64
		expected bool
	}
	testValues := []testValue{testValue{5.0, 5.1, 0.2, true}, testValue{19.1, 18.7, 0.1, false}}
	for i := 0; i < len(testValues); i++ {
		result := ApproxValue(testValues[i].a, testValues[i].b, testValues[i].epsilon)
		if result != testValues[i].expected {
			t.Errorf("I expected %t when testing ApproxValue(%f, %f, %f). I got %t", testValues[i].expected, testValues[i].a, testValues[i].b, testValues[i].epsilon, result)
		}
	}
}
