package main

import "fmt"

func main() {
	// Text type
	// String: unicode character
	// Strings can be treated as arrays, they are aliases for bytes, but immutable

	sss := "this is a string"
	// gives utf value default, have to get string value
	fmt.Printf("\n%v, %v,  %T\n", sss[2], string(sss[2]), sss[2])
	fmt.Println(sss + " which is cool")

	// ascii values
	bytes := []byte(sss)
	fmt.Printf("\n%v, %T\n", bytes, bytes)

	// Rune
	// UTF32
	// Any character can be upto 32b char
	// Rune are true type alias to int32
	r := 'a'
	fmt.Printf("\n%v, %T\n", r, r)

}
