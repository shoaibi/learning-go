package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	// receive only channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// send only channel
	go func(ch chan<- int) {
		i := 42
		ch <- i
		i = 43
		wg.Done()
	}(ch)
	wg.Wait()
}
