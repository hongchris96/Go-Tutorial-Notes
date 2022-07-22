package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

// Signal Only channel: zero memory allocation for sending message, but is able to let receiver know a message was sent
var doneCh = make(chan struct{}) // strongly typed to a struct with no field

func main() {
	// Channel
	fmt.Println("Channels:")
	// make function to create channels: keyword chan followed by data type that flows through the channel
	ch := make(chan int) // strongly typed

	wg.Add(2) // two goroutines here: channels are bidirectional by nature
	// but we want to set data flow restrictions by setting the goroutine params
	// receiving data from channel: receiving value <- channelName
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	// sending message into channel: channelName <- data
	go func(ch chan<- int) {
		// passing value into channel is passing a copy of the data
		ch <- 42
		// unbuffered channel, only one message can be in the channel at one time
		// need to have one receiver for each sender
		wg.Done()
	}(ch)
	wg.Wait()

	// adding buffers
	// second param of internal data store that can store 50 integers
	// if the sender or receiver operate at a different frequency than the other side
	chb := make(chan int, 50)

	wg.Add(2)
	// receiver
	go func(chb <-chan int) {
		// ranging over channel
		// for i := range chb {
		// 	fmt.Println(i)
		// }

		// processing when the channel is closed manually
		for {
			if i, ok := <-chb; ok { // ok tells you if channel is open or closed
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(chb)
	// sender
	go func(chb chan<- int) {
		chb <- 42
		chb <- 27  // no error cuz buffered
		close(chb) // closing the channel so that the receiver for loop knows when to end
		// can NOT pass message into a closed channel anymore
		wg.Done()
	}(chb)
	wg.Wait()

	// Select statements
	fmt.Println("Select statements:")
	// goroutine here monitor log channel for log entry coming in throughout application
	go logger()
	// defer func() {
	// 	close(logCh) // explicitly stating how goroutine ends, don't want it to be forcibly closed by main function end
	// }()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	// for entry := range logCh {
	// 	fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	// }

	for {
		select {
		// select statement here monitors the two channels
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
		// not putting default here, because we want select statement to block forever until message comes in
	}
}
