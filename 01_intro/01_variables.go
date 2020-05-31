package main

// Package names: singular and short. No common, util, shared, lib, ...

// Avoid using aliases imports unless there is _a conflict or last path of import doesn't match
import (
	"fmt"
)

func main() {
	// Verbose declaration
	var i int
	i = 42
	fmt.Println(i)

	// Semi verbose, with assignment directly
	// type can be omitted here
	var j int = 9
	fmt.Println(j)

	var k float32 = 22.1
	fmt.Printf("%v, %T", k, k)

	// This would never be float32 but float64
	l := 28.1
	fmt.Printf("\n%v, %T", l, l)

	fmt.Println("")
	// conversion:
	d := float32(9)
	fmt.Println(d)
}
