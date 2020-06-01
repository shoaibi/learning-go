package main

import (
	"fmt"
)

// structs are value types

// Generally better to replace naked booleans with constants
// Allows expansion of values in future
const (
	DoesNotPayTax = iota
	PaysTax
)

// same naming rules as the variable
type Student struct {
	name     string
	age      int
	grade    string
	subjects []string
	// this is the only field exposed to public
	TaxStatus int
}

func main() {
	// can skip field names here (positional syntax) and use positional but if we modify struct above
	// the assignments below would have to be adjusted
	// when using field names, the order doesn't have to be the same
	// enforced by `go vert` to use field names
	jason := Student{
		name:  "Jason",
		age:   20,
		grade: "A+",
		subjects: []string{
			"Math",
			"CS",
		},
	}

	fmt.Println(jason)
	fmt.Println(jason.age)
	fmt.Println(jason.subjects)
	fmt.Println(jason.subjects[1])
}
