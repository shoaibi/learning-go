package main

import (
	"fmt"
)

func main() {
	// short syntax
	// anonymous struct
	// good for dynamic data that is used quickly and discarded
	doctor := struct{ name string }{name: "John"}
	fmt.Println(doctor)

	// value types
	doctor2 := doctor
	doctor2.name = "Jason"
	fmt.Println(doctor2)
}
