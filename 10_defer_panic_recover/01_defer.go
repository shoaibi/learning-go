package main

import (
	"fmt"
	"log"
)

// defer cleanup code, closing connection, removing lock on mutex
// could also manually do cleanup at the end of function but
//     that might not get ccalled due to panics, returns, etc. and it isn't that readable either
// helps with adHD too
// Can use defer in conditions too

// defers are `link`ed chains of function points mimicking a stack
// defer(1.13): 35ns, function(): 3s. not that bad for performance unless it is a low latency system
// for low latency systems we could open-code it and put the defers
// explicitly by ourselves. 1.14 will do it this way but limits defers to 8/func, giving 4ns of call performance

type cleanupFunc func()
func createTestFiles() ([]string, cleanupFunc) {
	var testfiles []string
	// create files

	cleanupFunc := func() {
		// logic to cleanup those files
	}

	return testfiles, cleanupFunc
}

func main() {
	fmt.Println("start")
	// defer executed statement after the last instruction of the function
	// defer executes in LIFO order
	defer fmt.Println("second")
	defer fmt.Println("third")
	defer fmt.Println("fourth")
	fmt.Println("end")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// use defer to close the resource immediately after opening it
	// defer inside a loop is a bad choice as there would be way too many items to close

	a := "start"
	// Prints starts
	// Arguments are executed at the time of the defer call, not at the time of function execution
	defer fmt.Println(a)
	a = "end"


	// create a lot of test files
	testfiles, cleanup := createTestFiles()
	defer cleanup()

	firstDeferGotcha()

}

func firstDeferGotcha()  {
	// This would actually print five x5 because by the time defer is executed, item is already set to "give".
	for _, item := range []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	} {
		// Fix? pass item to the anonymous function
		defer func() {
			log.Println(item)
		}()
	}

}
