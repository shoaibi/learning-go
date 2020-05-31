package main

import "fmt"

func main() {

	switch i := 2 + 6; i {
	// no overlap allowed here
	case 1, 3, 5, 7:
		fmt.Println("Odd number till 7")
		// can still use `break` here to execute early if needed, else implied
	case 2, 4, 6, 8:
		fmt.Println("Even till 8")
	default:
		fmt.Println("Neither one or two")
	}

	// tagless switch
	value := 8
	switch {
	// overlaps are fine here
	case value <= 10:
		fmt.Println("Value is less than equal to 10")
		fallthrough // to make next case to also run, irrespective of its condition
	case value <= 15:
		fmt.Println("Value is less than equal to 15")
	case value <= 20:
		fmt.Println("Value is less than or equal to 20")
	default:
		fmt.Println("Greater than 20")
	}

	// Type switch
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is an stringg")
	case [2]int: // would only match array of size=2 and int array
		fmt.Println("i is an [2]int")
	default:
		fmt.Println("Unknown type")
	}
}
