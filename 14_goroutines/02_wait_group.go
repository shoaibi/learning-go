package main

import (
	"fmt"
	"sync"
)

// designed to  synchronize multiple goroutines together
var wg = sync.WaitGroup{}

func main() {
	g := "Message here"
	wg.Add(1)
	go func(g string) {
		fmt.Println(g)
		wg.Done()
	}(g)
	g = "Another"
	wg.Wait()

}
