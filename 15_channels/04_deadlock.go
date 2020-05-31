package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)
	wg.Add(1)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(j int) {
			ch <- j
			wg.Done()
		}(j)
	}
	wg.Wait()

	// deadlock because the outside goroutine would die after getting 0 but we have more
	// pushing goroutines inside loop
}
