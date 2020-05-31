package main

import "fmt"

type Incrementer interface {
	Increment() int
}

// Interface can be on any type where we can add methods
// Can't add it on int because it isn't in our package, actually it is a primitive type and we can't modify it
type IntCounter int

func (ic *IntCounter) Increment() int {
	// incrementing the type itself, which is an integer
	*ic++
	return int(*ic)
}

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}
}
