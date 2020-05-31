package main

import "fmt"

func main() {
	a := 42
	b := a
	b = 44
	// copied by value so both have different address
	fmt.Printf("%v, %v\n", a, &a)
	fmt.Printf("%v, %v\n", b, &b)
	fmt.Println()

	c := 40
	d := &c // or var d int* = &c
	*d = 41 // have to use * to tell to change the value
	// pointers, both should have same memory and same value
	fmt.Printf("%v, %v\n", c, &c)
	fmt.Printf("%v, %v\n", *d, d)
	fmt.Println()

	// pointer arithmetic isn't allowed
	// can't do &c + 8 to get the next memory address
	// actually can do it using unsafe package but not recommended
	f := [3]int{1, 2, 3}
	g := &f[0]
	h := &f[1]
	// h is 8 bytes higher than g, because arrays are contagious
	fmt.Printf("%v, %p, %p \n", f, g, h)
	// because it is a pointer over an array, tempted to do: %f[1]+8 to get next index but this is unsafe

}
