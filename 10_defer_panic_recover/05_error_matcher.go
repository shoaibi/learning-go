package main

import "fmt"

type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

//Exporting custom error types? Better to expose matcher functions
func IsNotFoundError(err error) bool {
	// type assertion. Always use comma,ok syntax to avoid panics for mismatches
	_, ok := err.(errNotFound)
	return ok
}
func Open(file string) error {
	return errNotFound{file: file}
}

func main() {
	// Usually this code would be in a different package
	if err := Open("foo"); err != nil {
		if IsNotFoundError(err) {
			fmt.Println("I got: ", err)
		} else {
			panic("unknown error")
		}
	}
}
