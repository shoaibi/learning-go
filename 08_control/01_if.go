package main

import "fmt"

func main() {
	// Never allowed to have a single line block. Must have curly braces
	if true {
		fmt.Println("The if condition is true")
	}

	grades := map[string]int{
		"A+": 90,
		"A":  80,
		"B":  70,
	}
	// initializer syntax
	// variables declared here are also available inside else block
	// recommended to use this reduced scope; unless it causes more nesting
	if popValue, ok := grades["B"]; ok {
		// popValue is only available in this block.
		fmt.Println(popValue)
	} else {
		fmt.Println("B is not set: ", popValue, ok)
	}
}
