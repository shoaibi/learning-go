package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 100)

	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		for j := 0; j < 10; j++ {
			ch <- j
			ch <- j * 10
		}

		// only close when sure, as a channel can't be reopened
		// we can't even detect if a channel is closed
		// better to have a defer to listen for panic and handle it there
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
