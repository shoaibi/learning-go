package main

import "fmt"

// Function names follow same style as variables
// First character:
//       Lower case = package
//       Capital case = public
// Test functions can have _ in the name, say MainFunction_TestScenarioName
// Functions are usually sorted in order of call and by receiver. Exported functions first.
func main() {
	fmt.Println("A function without any parameters and return value")
	name := "Jason"
	// use c-style comments to hint parameters
	sayHello("Welcome" /* greeting */ , name)
	fmt.Println(name)
	byReference(&name)
	fmt.Println(name)
}

// if parameters are of same type, we can group them
func sayHello(greetings, name string) {
	fmt.Println(greetings, name)
}

func byReference(name *string) {
	*name = "Jane"
}
