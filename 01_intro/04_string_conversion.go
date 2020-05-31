package main

// grouped import statement are called "factored"
import (
	"fmt"
	"strconv"
)

func main() {
	// This gives '*' because 42 is the ascii code for that
	s := string(42)
	fmt.Println(s)

	// string conversion:
	ss := strconv.Itoa(42)
	fmt.Println(ss)
}
