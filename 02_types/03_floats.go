package main

import (
	"fmt"
	"math"
)

func main() {
	// Floats
	// float32 and float64
	// Can use exponential notation, E or e.

	// Imaginary Numbers
	// Skipped

	f := 0.1
	fmt.Println(math.Pow(math.Sqrt(f), 2) == f)

	f = 0.123
	// here it is false
	// floating point numbers are approximation of decimal values
	fmt.Println(math.Pow(math.Sqrt(f), 2) == f)

	// workaround
	f = 0.123
	errThreshold := 0.001
	fmt.Println(math.Abs(f/math.Pow(math.Sqrt(f), 2)-1) < errThreshold)
}
