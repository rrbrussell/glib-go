package glib

// Clamp returns the value of x clamped between low and high
func Clamp(x float64, low float64, high float64) float64 {
	if x < low {
		return low
	}
	if x > high {
		return high
	}
	return x
}

// ApproxValue returns true if a and b are closer that epsilon from each other
func ApproxValue(a float64, b float64, epsilon float64) bool {
	if a < b {
		//log.Printf("%f is less than %f\n", a, b)
		result := b - a
		//log.Printf("%f-%f=%f\n", b, a, result)
		return result < epsilon
	}
	//log.Printf("%f is less than %f\n", b, a)
	result := a - b
	//log.Printf("%f-%f=%f\n", a, b, result)
	return result < epsilon
}
