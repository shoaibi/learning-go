package main

import "fmt"

func main() {
	// Boolean
	// It is its own type, not an alias
	var f bool = false
	fmt.Printf("%v, %T", f, f)

	t := true
	fmt.Printf("\n%v, %T\n", t, t)

	n := 1 > 2
	fmt.Printf("\n%v, %T\n", n, n)

	// Default is false
	var b bool
	fmt.Printf("\n%v, %T\n", b, b)
}
