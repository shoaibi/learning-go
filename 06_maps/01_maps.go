package main

import "fmt"

// maps are also by reference
// use copy() function after creating a new map to assign values
// careful when return them or reassigning them
// Use make where possible with capacity hint to allow the possibility of less allocations

func main() {
	// keys have to be tested for equality else they aren't valid, say maps or slices
	// slice has order, map doesn't
	// Passed by reference
	grades := map[string]int{
		"A": 90,
		"B": 80,
		"C": 70,
	}

	fmt.Println(grades)
	fmt.Println(len(grades))

	ages := make(map[string]int, 1)
	ages = map[string]int{
		"kid":      5,
		"boy":      10,
		"teenager": 14,
	}

	fmt.Println(ages)
	fmt.Println(ages["kid"])
	ages["adult"] = 20
	fmt.Println(ages)

	delete(ages, "boy")
	fmt.Println(ages)
	// it gives 0, because of int's default value?
	fmt.Println(ages["boy"])
	// check for existence
	_, ok := ages["boy"]
	fmt.Println(ok)
}
