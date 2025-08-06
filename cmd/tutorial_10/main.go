// Channels allow goroutines to communicate and synchronize
// by sending and receiving values.
//
// Unbuffered channels block the sender until the receiver receives the value.
//
// Buffered channels allow sending multiple values
// without blocking, up to the buffer capacity.
package main

import (
	"fmt"
	"time"
)

func main() {
	// unbufferedSample()
	// bufferedSample()
	// bufferedChannelSample()
	channelSample()
}

func bufferedSample() {
	fmt.Println("Buffered Channel Sample")
	// Initialize a Buffered channel with a capacity of 5
	// Buffered Channel Characteristics:
	// Sender and receiver can operate asynchronously (not at the same time).
	// Buffer → multiple values can be sent and received at a time.
	// Ensures safe synchronization but must be used carefully to avoid deadlocks.
	var channel = make(chan int, 5)
	// Start the concurrent process goroutine
	go process(channel)
	// Wait for the process to finish
	// And print the result directly from the channel
	for poppedValue := range channel {
		fmt.Println(poppedValue)
		time.Sleep(time.Second * 1)
	}
}

func unbufferedSample() {
	fmt.Println("Unbuffered Channel Sample")
	// Initialize an Unbuffered channel
	// Unbuffered Channel Characteristics:
	// Sender and receiver must be synchronized (operate at the same time).
	// No buffer → only one value can be sent and received at a time.
	// Ensures safe synchronization but must be used carefully to avoid deadlocks.
	var channel = make(chan int)
	// Start the concurrent process goroutine
	go process(channel)
	// Wait for the process to finish
	// And print the result directly from the channel
	for poppedValue := range channel {
		fmt.Println(poppedValue)
	}
	fmt.Println()
}

func process(channel chan int) {
	// Close the channel to signal the end of the data stream
	// So that the receiver knows there are no more values to receive
	// And the range loop in the main function can exit
	// Otherwise, the main function will hang indefinitely and we will get:
	// fatal error: all goroutines are asleep - deadlock!
	defer close(channel)

	// When this is used by an buffered channel
	// This won't wait for the receiver to pop the values
	// This will/might write every value to the channel way faster before the receiver pops them
	for i := range 5 {
		// Sets values to the channel
		channel <- i
	}

	fmt.Println("process function completed")
}
