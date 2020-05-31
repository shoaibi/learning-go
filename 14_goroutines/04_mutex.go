package main

import (
	"fmt"
	"runtime"
	"sync"
)

// safety comes at a cost
// do not embedded it in libraries, lets consumers decide

// embedding mutex in unexported structs is recommended
// for exported ones, add a as property
// mutex has a zero value so they can be used directly without new

// defer calls to unlock

// accessed globally but that is ok
var wg = sync.WaitGroup{}
var counter int = 0
var m sync.RWMutex

func main() {
	// Number of OS threads == Number of processors on a machine
	//runtime.GOMAXPROCS(1) single threaded
	// 1 OS thread / processor is minimum but we can optimize it. 100 makes more sense but any more would be overkill
	fmt.Printf("Available thread: %v, \n", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		// destroyed the parallelism
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
	// Bit messy and unpredictable behavior
}

func increment() {
	defer m.Unlock()
	counter++
	wg.Done()

}

func sayHello() {
	defer m.RUnlock()
	fmt.Printf("Hello with; %v\n", counter)
	wg.Done()
}
