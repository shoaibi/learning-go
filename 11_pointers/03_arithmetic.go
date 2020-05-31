package main

import "fmt"

func main() {
	// pointer arithmetic isn't allowed
	// can't do &g + 8 to get the next memory address
	// actually can do it using unsafe package but not recommended
	f := [3]int{1, 2, 3}
	g := &f[0]
	h := &f[1]
	// h is 8 bytes higher than g, because arrays are contagious
	fmt.Printf("%v, %p, %p \n", f, g, h)
	// because it is a pointer over an array, tempted to do: %f[1]+8 to get next index but this is unsafe
}
