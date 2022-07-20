package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutines
	fmt.Println("Goroutines:")

	// creating Goroutine: abstraction of a thread
	// scheduler maps goroutine onto OS threads for periods of time
	// cpu assign different goroutine processing time
	// goroutine have small stack spaces, can be relocated easily

	// spin off green thread, and run sayHello function in green thread
	// goroutine spawning off the main function
	go sayHello()

	var msg = "Hi"
	// anonymous go routine function
	go func() {
		fmt.Println(msg) // has access to the outer scope, but not best practice
	}()
	msg = "Goodbye" // lines in main function runs first, then the goroutine function

	time.Sleep(100 * time.Millisecond) // delay the main function return, so that goroutine function can run
	// Synchronization

	// Waitgroups

	// Mutexes

	// Parallelism

}

func sayHello() {
	fmt.Println("Hello")
}
