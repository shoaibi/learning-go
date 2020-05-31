package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)

	wg.Add(2)
	go func(ch <-chan int) {
		// deadlocking because we keep listening even when we have no more messages
		// time to close the channel
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 10
		ch <- 20
		wg.Done()
	}(ch)

	wg.Wait()

	/* for range loop is essentially doing this"
	for {
		if i, ok := <- ch; ok {
			fmt.Println(i)
		} else {
			break
		}
	}
	*/
}
