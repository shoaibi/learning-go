package main

import "fmt"

// everything can be cast to empty interface
// It has not interface
//    We can't call any methods
//    We'd need to do type conversion or do reflection
// Empty interface can also be defined explicitly as :
type emptyInterface interface{}

func main() {
	var a interface{}
	fmt.Println(a)

	var b emptyInterface
	fmt.Println(b)
}
