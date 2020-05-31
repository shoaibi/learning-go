package main

import "fmt"

// Always accept the most specific interface needed for that funcction/method

// why have smaller interfaces? more powerful
type Incrementer interface {
	Increment() int
}

type Decrementer interface {
	Decrement() int
}

type DoubleIncrementer interface {
	DoubleIncrement() int
}

// Composed of Incrementer and Decrementer by embedding both interfaces
type IncrementerDecrementer interface {
	Incrementer
	Decrementer
}

// MOAR COMPOSITION
type DoubleIncrementerDecrementer interface {
	IncrementerDecrementer
	DoubleIncrementer
}

// We could just use type myNumber int too.
type myNumber struct {
	value int
}

func (n *myNumber) Increment() int {
	// incrementing the type itself, which is an integer
	n.value++
	return n.value
}

func (n *myNumber) Decrement() int {
	// incrementing the type itself, which is an integer
	n.value--
	return n.value
}

func (n *myNumber) DoubleIncrement() int {
	// incrementing the type itself, which is an integer
	n.value++
	n.value++
	return n.value
}

// constructor method to help with setup of struct
func NewMyNumber(v int) *myNumber {
	return &myNumber{
		value: v,
	}
}

func main() {
	var icd DoubleIncrementerDecrementer = NewMyNumber(3)
	for i := 0; i < 5; i++ {
		fmt.Println(icd.DoubleIncrement())
		fmt.Println(icd.Decrement())
	}

	// Reverse type conversion
	// Now we can work with it as myNumber
	// Can only convert to types that implement the same interfaces
	var a = icd.(*myNumber)
	fmt.Println(a)
}
