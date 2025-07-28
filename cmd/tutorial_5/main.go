package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var myString = "résumé"
	var indexed = myString[1]
	fmt.Println(indexed)
	fmt.Printf("%v, %T\n\n\n", indexed, indexed)

	/*
		0 114
		1 233 (index 2 is skipped because é takes 2 bytes in utf-8)
		3 115
		4 117
		5 109
		6 233
	*/
	for i, val := range myString {
		fmt.Println(i, val)
	}

	fmt.Printf("The bytes of 'myString' is %v\n", len(myString))
	fmt.Printf("The real length of 'myString' is %v\n", utf8.RuneCountInString(myString))

	var myRunes []rune = []rune("résumé")
	fmt.Println("myRunes: ", myRunes)

	var myRune = 'a'
	fmt.Println("myRune: ", myRune)         // 97
	fmt.Println("myRune: ", string(myRune)) // Converts to string "a"

	// String building
	var strSlice = []string{"g", "e", "n", "e", "s"}
	var catStr = ""
	for i := range strSlice {
		// We are creating a new string everytime we concatenate
		// Which is pretty inefficient
		// Use go's built-in "strings" lib
		catStr += strSlice[i]
	}

	catStr += "sis"

	fmt.Println(catStr)

	// Strings are immutable
	// catStr[0] = "G" // cannot assign to catStr[0]
	// But you can do this instead, use string array or rune array and modify
	strSlice[0] = "M"

	// Efficient string builder
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Println("efficient strBuilder: ", strBuilder)
	// Only here "strBuilder.String()" creates a new string for the appended values
	fmt.Println("efficient strBuilder String: ", strBuilder.String())

}
