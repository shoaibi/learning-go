package main

import "fmt"

// constants can't be declared with :=
const a int16 = 32

func main() {
	// consts can't be set to a function call or anything that has to be computed at run time
	// `a` is shadowing the one from package level, we can also change type
	const a = "I am a constant"
	const b int = 42
	const c = 91

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)

	// This works because compile replaces all instances of constants with values
	var v int16 = 26
	fmt.Println(v + c)
}
