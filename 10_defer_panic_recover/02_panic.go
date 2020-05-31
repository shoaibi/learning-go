package main

import "fmt"

func main() {
	a, b := 1, 0
	// this gets printed before panic.
	// order: main, defer, panic, return value
	defer fmt.Println("this was deferred")
	ans := a / b
	fmt.Println(ans)
	// can also use panic("message")
}
