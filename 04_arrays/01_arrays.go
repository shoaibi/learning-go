package main

import "fmt"

func main() {
	// Fixed size of single type
	grades := [4]int{1, 5, 30, 20}
	fmt.Println(grades)

	// 3 different variables could be located anywhere
	// array data is contagious so access is faster

	// declaring array without saying the limit explicitly
	values := [...]int{5, 4}
	fmt.Println(values)

	// 4 elements array, which is empty
	var taxes [4]string
	fmt.Println(taxes)

	// Why use index? Tell compiler to walk (index * memoryRequiredForType) blocks forward
	taxes[0] = "Sales"
	fmt.Println(taxes)
	taxes[1] = "Withholding"
	fmt.Println(taxes)
	taxes[2] = "Pension"
	fmt.Println(taxes)
	taxes[3] = "Insurance"
	fmt.Println(taxes)
	fmt.Println(len(taxes))

	// array of arrays
	var matrix [2][2]int
	matrix[0] = [2]int{0, 0}
	matrix[1] = [2]int{1, 1}
	fmt.Println(matrix)
	// same thing as above but in one line
	var secondMatrix = [2][2]int{{2, 2}, {3, 3}}
	fmt.Println(secondMatrix)

	a := [...]int{1, 2, 3}
	b := a  // a new copy is assigned to b
	c := &a // reference
	b[1] = 8
	c[1] = 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
