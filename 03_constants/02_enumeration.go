package main

import "fmt"

// Enumerate constant
const (
	// iota is scoped at a const block
	io1 = iota
	io2 = iota
	// iota is auto assigned
	io3
	io4
	io5
)

// iota resets here
const (
	io6 = iota
	io7
)

func main() {
	// Enumeration
	fmt.Printf("%v\n", io1)
	fmt.Printf("%v\n", io2)
	fmt.Printf("%v\n", io3)
	fmt.Printf("%v\n", io4)
	fmt.Printf("%v\n", io5)
	// Enumeration resets
	fmt.Printf("%v\n", io6)
	fmt.Printf("%v\n", io7)
}
