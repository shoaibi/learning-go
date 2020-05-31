package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function")
	}()

	// better to pass the variable to avoid async issues
	// i would be available inside anonymous function though
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// each closure is bound to its own referenced variables
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
