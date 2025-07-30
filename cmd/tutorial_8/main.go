package main

import "fmt"

// Pointer is a special type that basically stores an address(memory locations)
// Pointer Variable
// var p = 0x140000b2020
// Normal boring variable
// var n = 1071

// Pass by value
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 is %p: ", &thing2)
	for i := range thing2 {
		thing2[i] *= thing2[i]
	}
	return thing2
}

// Pass by reference
func squareWithPassByReference(thing3 *[5]float64) *[5]float64 {
	fmt.Printf("The thing3 address is %p PassByReference: ", thing3)
	for i := range thing3 {
		thing3[i] *= thing3[i]
	}
	return thing3
}

func main() {
	var p *int32
	fmt.Println("pointer p: ", p)
	// Allocates a memory addres of type int32
	// If we don't assign this, we will get nil invalid memory address error!
	p = new(int32)
	fmt.Println("pointer p after: ", p)
	var i int32
	// Dereferencing the pointer
	// *p shows the "value" stored at that memory location
	// p -> 0x1400000e09c (Shows the address)
	// *p -> 0 (Shows the value at that memory location)
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v", i)
	*p = 1453
	fmt.Printf("The new value p points to is: %v\n", *p)

	// &i means the memory address of the variable, not its value
	// Here the address of pointer has changed
	// P refers to the memory address of i
	p = &i
	fmt.Printf("The new address of p is: %v\n", p)
	// i = 123 // Or this will also change the *p value
	*p = 1071
	fmt.Printf("The new value p points to is: %v\n", *p)

	// Slices
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice // Shallow copy, references are copied
	sliceCopy[2] = 4
	// Both will change here, because with slices
	// We are just copying the pointers!
	// So = slice will not be a deep copy
	fmt.Println("slice: ", slice)
	fmt.Println("sliceCopy: ", sliceCopy)

	// Deep copy
	original := []int{1, 2, 3, 4, 5}
	deepCopy := make([]int, len(original))
	copy(deepCopy, original)

	// Or just one line
	// deepCopy := append([]int(nil), original...)
	fmt.Println("original: ", original)
	fmt.Println("deepCopy: ", deepCopy)

	stringSlice := []string{"A", "BSAD", "GEN"}
	stringSlice = append(stringSlice, "GENESIS")
	fmt.Println(stringSlice)

	// var thing1 = [5]float64{1, 2, 3, 4, 5}
	// var result = square(thing1)
	// fmt.Println("thing1: ", thing1)
	// fmt.Println("result: ", result)
	// fmt.Printf("The memory location of the thing1 is %p: \n", &thing1)
	// fmt.Printf("The memory location of the result is %p: \n", &result)

	var thing3 = [5]float64{1, 2, 3, 4, 5}
	var result3 = squareWithPassByReference(&thing3)
	fmt.Println("result2: ", result3)
	fmt.Printf("\nThe memory location of the thing3 is %p: ", &thing3)
	fmt.Printf("\nThe memory location of the result3 is %p: ", result3)
}
