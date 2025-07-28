package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Genesis Corp.")

	var numerator int = 7
	var denominator int = 1
	result, remainder, err := intDivision(numerator, denominator)
	println("remainder: ", remainder, "\n\n")
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result is %v with reaminder\n", result)
	} else {
		fmt.Printf("Result is %v with remainder %v\n", result, remainder)
	}

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("Result is %v with reaminder\n", result)
	default:
		fmt.Println("Default")
	}

	if 1 == 1 && 2 == 2 || true {
		fmt.Println(true)
	}

	switch remainder {
	case 0:
		fmt.Println("remainder: ", remainder)
	// 1 or 2
	case 1, 2:
		fmt.Println("remainder 1 or 2: ", remainder)
	default:
		fmt.Println("default")
	}
}

func printMe(printValue string) {
	fmt.Println("Func: ", printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
