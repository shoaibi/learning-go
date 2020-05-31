package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// specify buffer length
	ch := make(chan int, 20)

	// buffer channels allows sender and receiver to run at different rates
	// say when dealing with burst traffic

	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		// but this is ugly. We have to manually get the second message
		// better to use for-range loops
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 10
		ch <- 20
		wg.Done()
	}(ch)

	wg.Wait()
}
