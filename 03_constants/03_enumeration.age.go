package main

import "fmt"

// this supports also an initial offset, basic math operations
const (
	child = iota
	boy
	teenager
	adult
	middleAge
	oldMan
)

func main() {
	// Use of enumerated constants
	var user = boy
	fmt.Printf("%v\n", user == boy)
	// But there is a problem.
	// Default value of integer is 0 which is same as child above
	// Too work around; use the zero value has an error value or use underscore symbol to throw it away
	var me int
	fmt.Printf("%v\n", me == child)

	// though in some cases it does make sense;
	/*
		type LogOutput int
		const (
			LogToStdout LogOutput = iota
			LogToFile
			LogToRemote
		)
	*/
}
