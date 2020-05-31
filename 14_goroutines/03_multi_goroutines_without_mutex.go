package main

import (
	"fmt"
	"sync"
)

// accessed globally but that is ok
var wg = sync.WaitGroup{}
var counter int = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
	// Bit messy and unpredictable behavior
}

func increment() {
	counter++
	wg.Done()
}

func sayHello() {
	fmt.Printf("Hello with; %v\n", counter)
	wg.Done()
}
