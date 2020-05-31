package main

import (
	"fmt"
)

// Use godoc fmt Println to get documentation on cli

// outside functions entities must have var. We can't use := here

// prefix unexported global declarations of var and consts with _, except for errors which start with err

// factored declaration
var (
	_a = 1
	_b = 2
	_c = "The Sea"
)

func main() {
	// Block scope starts

	fmt.Println(_a)
	fmt.Println(_b)
	fmt.Println(_c)

	// shadowed, inner scope takes precedence
	a := 22
	fmt.Println(a)
	fmt.Println(_b)
	fmt.Println(_c)
}
