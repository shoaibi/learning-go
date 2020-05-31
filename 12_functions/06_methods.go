package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

// prefer functions over methods, unless you need to modify the state
// function take input and given same input they give same output
// method usually modify internal

// Pointer vs value receiver
// Be consistent
// Small structure? Value receiver would be cheaper
// Large structure? Pointer receiver would be cheaper
// Pointer ones aren't safe for concurrency, sshared state

// functions with pointer receiver must get a pointer
// function with value receiver must get a value
// methods with pointer receiver can work with both, point and value
// methods with value receiver can also get a pointer

// Value receiver, receives full object
func (g greeter) greet() {
	// g is copy here, changes wouldn't go back to passed object
	g.name = "New name in greey"
	fmt.Println(g.greeting, g.name)
}

// Pointer receiver, receives pointer to actual object
func (g *greeter) pointerGreet() {
	// got pointer, any changes made here would appear in made struct
	g.name = g.name + " modified name"

	// can also use (*g).greeting and (*g).name due to implicit de-referencing of pointers
	fmt.Println(g.greeting, g.name)
}

func main() {
	g := greeter{
		greeting: "Welcome back",
		name:     "dear Jessie",
	}

	g.greet()
	fmt.Println(g)
	fmt.Println()
	g.pointerGreet()
	fmt.Println(g)
}
