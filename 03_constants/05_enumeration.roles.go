package main

import "fmt"

// roles, 2, 4, 8, 16, ...
const (
	isAdmin = 1 << iota
	isInHeadquarters
	canSeeFinancials
	canSeeUsers
	canSeeBookings
	canSeeBills
	canChangeUsers
)

func main() {
	// compression 3 roles into 8 bytes of data
	var roles byte = isAdmin | canSeeBills | canSeeBookings
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("can change users? %v\n", canChangeUsers&roles == canChangeUsers)
}
