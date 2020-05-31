package main

import "fmt"

func main() {
	// can't call f() before it is defined because it is defined as variable
	// longer version: var f func() = func() {
	f := func() {
		fmt.Println("Function as a type")
	}
	f()

	divFunc := getDividerFunction()
	fmt.Println(divFunc(6, 2))
	fmt.Println(divFunc(6, 0))
}

// Closures
// Can have a function that returns another function
// The return value have to be hinted properly
// For the return types we just need to define types
func getDividerFunction() func(a, b float64) (float64, error) {
	f := func(a, b float64) (result float64, err error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Can not divide by zero: %v/%v ", a, b)
		}

		return a / b, nil
	}
	return f

}
