package main

import (
	"fmt"
)

const (
	CanNotFly = iota
	CanFly
)

// inheritance? Embedding/Composition e.g. has-A relationship
// Embedding is not right for modeling behavior, Embedding allows behaviors to carry through
// Use interface to use types interchangeable
// Here the types aren't interchangeable
type Animal struct {
	Name   string
	Origin string
}

// Bird has-a Animal, has Animal like characteristics
type Bird struct {
	// notice the missing type of "Animal" here
	// If we add the type it is not embedding,
	//    it'd make Animal a property of Animal type
	// Embedded types to appear at top of stucts, followed by space
	Animal

	SpeedKPH float32
	AbleToFly   int
}

func main() {
	// In literal notation, developer must know that we have a nested type
	emu := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		AbleToFly:   CanNotFly,
	}
	fmt.Println(emu)

	// Here we don't need to know about embedded property
	ostrich := Bird{}
	ostrich.Name = "Ostrich"
	ostrich.Origin = "SSomewhere"
	ostrich.SpeedKPH = 30
	ostrich.AbleToFly = CanNotFly
	fmt.Println(ostrich)
}

// Avoid Embedding on exported structs. Use a private interface (pointer) property and
//     provide proxy methods to delegate. Tedious but encapsulation and abstraction.
// calls to underlying property to have a consistent API
// Embedding gives the outer type gets implicit copies of the embedded type's methods. We also get a
//    a field by the same name as the type. This means that if the embedded type is public,
//    the field is public. To keep backwards compatibility the putter type must expose this embedded type in future.
//    Embedding is convenience that helps you avoid writing tedious delegate methods.
// Embedding adds challenges around breaking changes
//  - Adding/Removing methods to embedded type
//  - Removing embedded type
//  - Replacing embedded type with another that satisfies the same interface
