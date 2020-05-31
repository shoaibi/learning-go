package main

import "fmt"

func main() {
	// slices have variable size
	// arrays back slices
	// slices are naturally reference types
	a := []int{1, 2, 3}
	printSlice("a", a)
	fmt.Println()

	b := a
	printSlice("b", b)
	fmt.Println("Changing first index in b to 99")
	b[1] = 999
	printSlice("a", a)
	printSlice("b", b)
	fmt.Println()

	// slicing options (<3 to Python)
	// range: start: inclusive, end:exclusive
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := a[:]   // all elements
	d := a[3:]  // elements from 3rd index onwards
	e := a[:6]  // elements till 6th index (exclusive)
	f := a[3:6] // elements from 3rd index inclusive till 6th index exclusive
	printSlice("a", a)
	printSlice("c", c)
	printSlice("d", d)
	printSlice("e", e)
	printSlice("f", f)

	// slices are by reference so all of the slices get updated
	fmt.Println("Changing a's 5th element to 99")
	a[5] = 99
	printSlice("a", a)
	printSlice("c", c)
	printSlice("d", d)
	printSlice("e", e)
	printSlice("f", f)
	fmt.Println()

	s := make([]int, 3)
	printSlice("s", s)
	fmt.Println()

	// pre-provision capacity to avoid copying underlying array
	ss := make([]int, 3, 10)
	printSlice("ss", ss)
	// append and capacity doubles
	fmt.Println("Appending 10, 20, 30")
	ss = append(ss, 10, 20, 30)
	printSlice("ss", ss)

	fmt.Println("Trying to go beyond capacity, adding 40-110")
	// going beyond the planned capacity doubles capacity against last known capacity
	ss = append(ss, 40, 50, 60, 70, 80, 90, 100, 110)
	printSlice("ss", ss)

	// concatenating 2 slices
	// slice size doubles from last size, not against the capacity
	fmt.Println("Concatenating 2 slices using expansion operator")
	ss = append(ss, []int{120, 130, 150, 160, 170, 180, 190, 200}...)
	printSlice("ss", ss)

	// remove from start of a slice
	printSlice("ss:", ss)
	b = ss[1:]
	printSlice("removing first:", b)
	// remove from end
	c = ss[:len(ss)-1]
	printSlice("Removing last:", c)
	// removing from middle; doesn't work, need to use a loop to create a copy of slice of ss
	d = append(ss[:2], ss[4:7]...)
	printSlice("Removing from middle:", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
