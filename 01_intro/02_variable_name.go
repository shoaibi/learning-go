package main

import (
	"fmt"
)

// First letter: Lower case variables are meant for Package; package level; all src files in same package can access
// First letter: upper  case variables are exposed to outside
// No private scope; use block scope instead

// Length of variable name should be based on lifespan of the variable
// Variables that have longer life span should have longer names
// Package level variables should have verbose names
// Acronyms should be all upper case.

func main() {
	fmt.Println("Hello", "World")
}
