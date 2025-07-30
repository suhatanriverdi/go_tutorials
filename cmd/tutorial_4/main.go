package main

import (
	"fmt"
	"time"
)

const LEN int = 1000000

// Benefits of setting the capacity of your slice ahead of time
func getPreAllocatedSlice() []int {
	// We set the capacity beforehand here
	// It basically means a space allocation beforehand
	var slice = make([]int, 0, LEN)
	for len(slice) < LEN {
		slice = append(slice, 0)
	}
	return slice
}

// We don't set any capacity here
func getNOTPreAllocatedSlice() []int {
	var slice = []int{}
	for len(slice) < LEN {
		slice = append(slice, 0)
	}
	return slice
}

func testPreallocatedSlices() (time.Duration, time.Duration) {
	var t0 = time.Now()
	getPreAllocatedSlice()
	var time0 = time.Since(t0)

	var t1 = time.Now()
	getNOTPreAllocatedSlice()
	var time1 = time.Since(t1)

	return time0, time1
}

func main() {
	var intArr [3]int32
	intArr[0] = 1453
	fmt.Println("intArr: ", intArr)
	fmt.Println("intArr: ", intArr[0])
	fmt.Println("intArr: ", intArr[1:3])

	fmt.Println("&intArr[0]: ", &intArr[0])
	fmt.Println("&intArr[1]: ", &intArr[1])
	fmt.Println("&intArr[2]: ", &intArr[2])

	var secondArr [5]int32 = [5]int32{1, 2, 3, 4, 5}
	// Alternative ways
	// var secondArr = [5]int32{1, 2, 3, 4, 5}
	// var secondArr = [...]int32{1, 2, 3, 4, 5}
	integerArr := [2]int32{}
	fmt.Println("integerArr: ", integerArr)
	fmt.Println("secondArr: ", secondArr)

	// Slices are flexible arrays
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Printf("The length is %v with capacity %v, \n", len(intSlice), cap(intSlice))
	fmt.Println("intSlice: ", intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println("After append slice: ", intSlice)
	fmt.Printf("After The length is %v with capacity %v, \n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println("After multiple appends slice: ", intSlice)

	// Alternatively
	var intSlice3 []int32 = make([]int32, 3)
	fmt.Println("Alternative slice: ", intSlice3, "\n\n")

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println("myMap: ", myMap)

	// Alternatively
	var myMap2 map[string]uint8 = map[string]uint8{"Melo": 23, "Genesis": 145}
	fmt.Println("myMap2: ", myMap2)
	fmt.Println("Melo: ", myMap2["Melo"])

	// Key does not exists but Go will return 0!
	fmt.Println("Melo: ", myMap2["NO"])
	fmt.Println("myMap2: ", myMap2)

	var age, has = myMap2["Adam"]
	fmt.Println("age: ", age, " has: ", has)
	if has {
		fmt.Println("The age is : ", age)
	} else {
		fmt.Println("Key not found!")
	}

	// Deletion from map
	// No return value is given
	delete(myMap2, "Adam")

	// No order is preserved on maps
	fmt.Println("Iterating over map: ")
	for key, value := range myMap2 {
		fmt.Println(key, value)
	}

	var intArr2 []int32 = []int32{5, 6, 7}
	for index, value := range intArr2 {
		fmt.Println(index, value)
	}

	// Go does not have while loop, we use for loop again
	// While loop sample
	var i int = 0
	for i < 10 {
		fmt.Print(i, " ")
		i = i + 1
	}

	// Alternative way
	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Print(i, " ")
		i = i + 1
	}

	fmt.Println(" ")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	for i := range 10 {
		fmt.Print(i, " ")
	}

	// Test Preallocating the Capacity
	var t0, t1 = testPreallocatedSlices()
	println("\nPreallocation  enabled: ", t0.String())
	println("Preallocation disabled: ", t1.String())
}
