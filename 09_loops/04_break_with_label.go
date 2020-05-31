package main

import "fmt"

func main() {
	// Breaking out with labels
Loop:
	for k := 0; k < 10; k++ {
		for l := 0; l < 10; l++ {
			if k%2 != 0 {
				continue
			}

			if k*l >= 5 {
				fmt.Println("k = ", k)
				fmt.Println("l = ", l)
				break Loop // need label so we can break upto that level
			}
		}
	}
}
