package main

import "fmt"

func main() {
	// i is scoped to for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// can't increment k next to j here as j++ is also a statement
	for j, k := 0, 0; j < 10; j++ {
		fmt.Println("j = ", j)
		fmt.Println("k = ", k)
		j += 2
	}
}
