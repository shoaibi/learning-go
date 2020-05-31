package main

import "fmt"

func main() {
	// Integer types
	// int8 (-128 - 127), int 16(-32768 - 32767), int32, int64
	// unsigned type for every signed one.
	// byte is unint8
	// can't add two integers of different types

	uninetTest := 322
	fmt.Println(uint32(uninetTest))

	// do not use integers for time
	// use time.Time and its functions
	// Use time.Duration for periods of time
}
