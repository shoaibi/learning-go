package main

import (
	"fmt"
	"time"
)

// main function also runs in a go routine
func main() {
	// golang creates a green thread
	// Traditional programming languages create OS thread. Individual function call stack for each thread.
	//     Pretty big. 1MB of RAM, time to setup. We try to avoid
	//        creating new threads and go into concept of thread pools
	// Follows Erlang to create a green threads.
	//    Abstract of a thread.
	//    Schedules go routines to already created OS threads for periods of time
	//    Can start with small stack space, cheaper to create an destroy
	// main function doesn't run it, go routine does but it needs time to print so we need to add
	// sleep so the go routine can finish.
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	// works because closures exists in go too
	// anonymous function does have access to outer variable even though running in go routine
	// but we created a dependency here between the go routine and the variable
	// and that is bad because the variable could change
	message := "Message here"
	go func() {
		fmt.Println(message)
	}()
	message = "Another"
	// Scheduler created the goroutine above but only gives it love once it sees sleep
	// The idea is to not interrupt the main thread
	// creates a race condition
	time.Sleep(100 * time.Millisecond)

	// better version; push it into function as arguments
	// pass by value.
	g := "Message here"
	go func(g string) {
		fmt.Println(g)
	}(g)
	g = "Another"
	time.Sleep(100 * time.Millisecond)

	// but still bad because we are sleeping.
	// Use waitgroups from synchronization

}

func sayHello() {
	fmt.Println("Hello")
}
