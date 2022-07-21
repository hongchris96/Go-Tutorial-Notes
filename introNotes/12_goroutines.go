package main

import (
	"fmt"
	"runtime"
	"sync"
)

// waitgroup
var wgroup = sync.WaitGroup{}
var counter = 0

// mutex: lock that application is going to honor
// protect parts of code, so only one entity can manipulate the code at a time
var m = sync.RWMutex{} // RWMutex only one can write at a time, but many can read

func main() {
	// Goroutines
	fmt.Println("Goroutines:")
	// concurrent programming

	// creating Goroutine: abstraction of a thread
	// scheduler maps goroutine onto OS threads for periods of time
	// cpu assign different goroutine processing time
	// goroutine have small stack spaces, can be relocated easily

	// spin off green thread, and run sayHello function in green thread
	// goroutine spawning off the main function
	// go sayHello()

	var msg = "Hi"
	// Waitgroups
	wgroup.Add(1) // like async await
	// anonymous go routine function
	go func(msg string) {
		fmt.Println(msg) // has access to the outer scope, but not best practice cuz of race condition
		wgroup.Done()
	}(msg) // pass in value as args is better
	msg = "Goodbye" // lines in main function runs first, then the goroutine function
	wgroup.Wait()
	// time.Sleep(100 * time.Millisecond) // delay the main function return, so that goroutine function can run, not best practice

	// waitgroup example 2:
	fmt.Println("Unordered")
	for i := 0; i < 10; i++ {
		// no synchronization here: race condition, unreliable order
		wgroup.Add(2)
		go sayHello()
		go increment()
	}
	wgroup.Wait()

	// Mutexes
	counter = 0
	fmt.Println("Mutex: ordered")
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wgroup.Add(2)
		m.RLock() // a read lock, before the goroutine executes
		go sayHello2()
		m.Lock() // a write lock
		go increment2()
	}
	wgroup.Wait()

	// what is gomaxprocs
	// by default shows the number of OS threads based on number of cores on the machine
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) // -1 number means the previous value of threads
	// number here can be changed to set number of threads to run this app

	// BEST practices:
	// - Don't create goroutines in libraries: let consumer control the concurrency
	// - Create goroutine knowing how it will end: avoid memory leaks
	// - Check for race conditions at compile time: go run -race src/main.go

}

func sayHello() {
	fmt.Printf("Hello %v\n", counter)
	wgroup.Done()
}

func increment() {
	counter++
	wgroup.Done()
}

func sayHello2() {
	fmt.Printf("Hello %v\n", counter) // protected
	m.RUnlock()                       // unlock when go routine is done
	wgroup.Done()
}

func increment2() {
	counter++ // protected
	m.Unlock()
	wgroup.Done()
}
