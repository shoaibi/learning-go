package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// channel is strongly typed
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		i := 42
		// data is sent here, e.g. 42
		// this line blocks the execution as channel is unbuffered
		//    so it will wait till message gets processed
		ch <- i
		i = 43
		wg.Done()
	}()
	wg.Wait()
}
