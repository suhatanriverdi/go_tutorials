package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 5

var waitGroup = sync.WaitGroup{}

func bufferedChannelSample() {
	fmt.Println("Buffered Channel Real Life Example")

	var chickenChannel = make(chan []string, 3)

	var websites = []string{"bim.com", "a101.com", "migros.com"}

	// Start checking chicken prices for each website
	for _, website := range websites {
		waitGroup.Add(1)
		go checkChickenPrices(website, chickenChannel)
	}

	// Prevent deadlocks by ensuring all goroutines finish before closing the channel
	go func() {
		// This is blocking goroutine that waits for all goroutines to finish and closes the channel
		waitGroup.Wait()
		close(chickenChannel)
	}()

	// You have to close the channel otherwise range will wait forever
	for poppedValue := range chickenChannel {
		fmt.Println(poppedValue[0], poppedValue[1], poppedValue[2])
	}
}

func checkChickenPrices(website string, chickenChannel chan []string) {
	var delay = time.Duration(rand.Intn(3000)) * time.Millisecond
	time.Sleep(delay)

	var chickenPrice float32
	for {
		chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- []string{
				website,
				fmt.Sprintf("Price: %.2f$", chickenPrice),
				fmt.Sprintf("Waited %.2f seconds", delay.Seconds())}
			break
		}
	}

	waitGroup.Done()
}
