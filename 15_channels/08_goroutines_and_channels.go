package main

import (
	"log"
	"math/rand"
	"sync"
)

const mapSize = 5

func main() {
	rand.Seed(42)

	// The idea behind this program is to have 3 parallel processes
	// produce, consumer and digest
	// so we have a pipeline of tasks that are doing across channels in this way:
	// produce -> transformChannel -> consume -> digestChannel -> digest
	// I could use wg.Wait() after each of these but that will make the process
	// more serial and I am trying to avoid that.
	// In terms of production code this could be a scenario where each
	// entity has to be processed and inserted immediately instead
	// of batch pull, batch consume and batch digest

	transformChan := produce()
	digestChan := consume(transformChan)
	data := digest(digestChan)

	log.Println(data)
}

func produce() chan int {
	// intentionally reduced size channel to
	// ensure consumption happens in parallel
	transformChan := make(chan int, mapSize-3)

	var wg sync.WaitGroup
	wg.Add(mapSize)

	for i := 0; i < mapSize; i++ {
		log.Println("[produce] Pushing data to channel")
		// Assume this is an expensive and lengthy operation
		go func() {
			defer wg.Done()
			transformChan <- getNumber()
		}()
	}

	// a separate go routine to not block main thread
	// close the channel when we are done with the main go routine above
	go func() {
		wg.Wait()
		close(transformChan)
	}()

	return transformChan
}

func getNumber() int {
	val := rand.Intn(100)
	log.Println("[push] Sending: ", val)
	return val
}

func consume(transformChan chan int) chan int {
	// intentionally reduced size channel to
	// ensure consumption happens in parallel
	var digestChan = make(chan int, mapSize-3)

	// Go routines work while the program is alive, doesn't have much to do with
	// the caller function except the variables
	go func() {
		defer close(digestChan)
		for i := range transformChan {
			log.Println("[consume] Got: ", i)
			// purely academic, production code would have some form of transformations
			// happening jere
			digestChan <- i * 10 * mapSize
		}
	}()

	return digestChan
}

func digest(digestChan chan int) map[int]int {
	data := make(map[int]int, mapSize)

	for i := range digestChan {
		log.Println("[digest] Got: ", i)
		// purely academic, production code would have some form of transformations
		// happening jere
		data[i/10/mapSize] = i
	}

	return data
}
