package main

import "fmt"

type myStruct struct {
	value int
}

func main() {
	ol := &myStruct{value: 9}
	fmt.Println(ol)

	// zero value for pointer is <nil>
	// useful to avoid passing whole structs around
	var ms *myStruct
	fmt.Println(ms)
	// recommended to  use &myStruct{} instead
	ms = new(myStruct) // can't specify properties in new()
	fmt.Println(ms)

	// (*ms) because * has lower precendece than .
	(*ms).value = 42
	fmt.Println((*ms).value)

	// compiler's syntactical sugar, same as above but easier to type
	ms.value = 44
	fmt.Println(ms.value)

	// slice's internal representation is a pointer to array
	// same is true for maps

}
