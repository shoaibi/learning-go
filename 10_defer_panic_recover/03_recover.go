package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start: main")
	panicker()
	fmt.Println("end: main")
}

func panicker() {
	fmt.Println("start: panicker")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error in panicker: ", err)
			// can re-panic here if we can't handle the error
		}
	}()
	defer fmt.Println("deferred middle: panicker")
	panic("I am panicking")
	// stops execution here but the higher order functions still keep working
	fmt.Println("end: panicker")
}
