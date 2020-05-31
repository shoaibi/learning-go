package main

import "fmt"

func main() {
	fmt.Println(returnByValue())
	fmt.Println(*returnByPointer())
	fmt.Println(returnWithoutName())
	fmt.Println(divideNumbers(5, 3))
	fmt.Println(divideNumbers(5, 0))
}

func returnByValue() int {
	return 311
}

// we can also return by reference
// return value would be promoted from function stack to heap when returning by reference
func returnByPointer() *int {
	sum := 2000
	return &sum
}

// not recommend; but confusing to read
// naked return
func returnWithoutName() (sum int) {
	// sum is already declared due to function signature
	sum = 3333
	return
}

// multiple returns are recommend for raising errors
// For the return types we just need to define types, we can add names but they're not used unless we do an empty return
func divideNumbers(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Can not divide by zero: %v/%v ", a, b)
	}
	return a / b, nil

}
