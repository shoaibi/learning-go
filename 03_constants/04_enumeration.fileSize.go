package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// trailing dot to make it float
	fileSize := 20000000000.
	fmt.Printf("%.2f GB\n", fileSize/GB)
}
