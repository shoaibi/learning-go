package main

import (
	"fmt"
	"time"
)

const (
	lvlInfo    = "INFO"
	lvlWarning = "WARNING"
	lvlError   = "EEROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logChan = make(chan logEntry, 50)
var doneChan = make(chan struct{}) // could use boolean but then a variable need to be allocated and copied

func main() {
	go logger()
	logChan <- logEntry{
		time:     time.Now(),
		severity: lvlInfo,
		message:  "Application Start",
	}

	logChan <- logEntry{
		time:     time.Now(),
		severity: lvlInfo,
		message:  "Application Shutdown",
	}

	// go routine: logger, ends abruptly
	// few ways to handle it:

	// we use a defer here to close the channel to logger subroutine gets that and closes
	//defer func() {close(logChan)}()

	// we use another channel to indicate that we are done, then use select in logger goroutine to break
	doneChan <- struct{}{}

}

func logger() {
	for {
		select {
		// if multiple channels get the message at the same time then behavior is undefined
		case log := <-logChan:
			fmt.Printf("%v - [%v] %v\n", log.time.Format("2006-01-02T15:04:05"), log.severity, log.message)
		case <-doneChan:
			break
		// can have a default statement, it won't be a blocking select statement
		// without default, select is blocking till it gets a message on any change
		default:

		}
	}
}
