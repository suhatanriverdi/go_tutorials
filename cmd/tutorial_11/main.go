package main

import (
	"fmt"
)

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func main() {
	var intSlice = []int{1, 2, 3}
	// Alternatively you can omit the type parameter
	// fmt.Println(sumSlice(intSlice))
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(isEmpty[int](intSlice))

	var floatSlice = []float64{1.1, 2.2, 3.3}
	fmt.Println(sumSlice[float64](floatSlice))
	fmt.Println(isEmpty[float64](floatSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3}
	fmt.Println(sumSlice[float32](float32Slice))
	fmt.Println(isEmpty[float32](float32Slice))

	// Use the jsonExample function
	jsonExample()

	// Use the engineExample function
	engineExample()
}
