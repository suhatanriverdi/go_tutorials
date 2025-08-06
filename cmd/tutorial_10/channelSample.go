package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE1 float32 = 5
const MAX_TOFU_PRICE1 float32 = 3

func channelSample() {
	fmt.Println("Channel Real Life Example")
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"bim.com", "a101.com", "migros.com"}
	for _, website := range websites {
		go checkChickenPricesA(website, chickenChannel)
		go checkTofuPrices(website, tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		var delay = time.Second * 1
		time.Sleep(delay)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE1 {
			tofuChannel <- website
		}
	}
}

func checkChickenPricesA(website string, chickenChannel chan string) {
	for {
		var delay = time.Second * 1
		time.Sleep(delay)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE1 {
			chickenChannel <- website
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// fmt.Printf("Found a deal on chicken at %s\n", <-chickenChannel)
	// fmt.Printf("Found a deal on tofu at %s\n", <-tofuChannel)
	//
	// This is like a switch/if statement but for channels
	// This select statement will listen to both channels and print a message based on the received message
	// And then print a message based on the received message and exit
	select {
	// If received a message from chickenChannel, print a text message
	// If received a message from tofuChannel, print an email message
	case website := <-chickenChannel:
		fmt.Printf("Text send, Found a deal on chicken at %s\n", website)
	case website := <-tofuChannel:
		fmt.Printf("Email send, Found a deal on tofu at %s\n", website)
	}
}
