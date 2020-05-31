package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}

// variadic parameter should always be the last one
func sum(values ...int) int {
	fmt.Println("Arguments", values)
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum
}
