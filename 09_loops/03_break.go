package main

import "fmt"

func main() {
	// condition within the loop
	i := 0
	for {

		if i%2 == 0 {
			fmt.Println("even i = ", i)
		} else {
			fmt.Println("odd i = ", i)
		}

		if i >= 9 {
			fmt.Println("i is >= 9; Breaking")
			break
		}
		i++
	}
}
