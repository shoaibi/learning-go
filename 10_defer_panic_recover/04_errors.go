package main

import "fmt"

/*
How to create errors?
- errors.New for errors with simple static strings
- fmt.Errorf for formatted error strings
- Custom types that implement an Error() method
- Wrapped errors using "pkg/errors".Wrap

Decision logic could be;
- Is this a simple error that needs no extra information? If so, errors.New should suffice.
- Do the clients need to detect and handle this error? If so, you should use a custom type, and implement the Error() method.
- Otherwise, fmt.Errorf is okay.
*/

type ErrorCode uint16

type Error struct {
	Code ErrorCode

	Message string
	Context map[string]string
}

func (e Error) Error() string {
	return fmt.Sprintf("[%v] %v", e.Code, e.Message)
}

func main() {
	e := Error{
		Code:    22,
		Message: "Error is 22",
	}

	fmt.Println(e)
	panic(e)
}
