package main

import "fmt"

func main() {
	// iterating over slices, arrays, maps or strings(unicode representation), channel
	ss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// could use _ for k if we don't want keys, can skip v directly
	for k, v := range ss {
		fmt.Println(k, v)
	}
}
