package main

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 1450
	intNum = intNum + 3
	fmt.Println(intNum)

	var floatNum float32 = 10.1
	fmt.Println(floatNum)
	var intNum2 int = 123
	var result float32 = floatNum + float32(intNum2)
	fmt.Println(result)

	var intNum1 int = 5
	var intNum3 int = 2
	fmt.Println("Operations: ")
	fmt.Println(intNum1 / intNum3)
	fmt.Println(intNum1 % intNum3)
	fmt.Println(intNum1 * intNum3)
	fmt.Println(intNum1 - intNum3)
	fmt.Println(intNum1 + intNum3)

	var genesisString string = "Genesis" + "Corp"
	var myString string = `Genesis
						   -Corp`
	fmt.Println(genesisString)
	fmt.Println(myString)

	fmt.Println("Num of bytes: ", len(myString))
	fmt.Println("length: ", utf8.RuneCountInString(genesisString))

	var myRune rune = 'a'
	fmt.Println(myRune, utf8.RuneLen(myRune))

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum5 int
	fmt.Println(intNum5)

	var testInferTypeFromUsage = "Test"
	fmt.Println(testInferTypeFromUsage, utf8.RuneCountInString(testInferTypeFromUsage))

	anyVariable := "genesis-corp"
	fmt.Println(anyVariable)

	var var12, var21 = 1, 2
	var var1, var2 int = 1, 2
	var3, var4 := 1, 2
	fmt.Println(var1, var2)
	fmt.Println(var3, var4)
	fmt.Println(var12, var21)

	const constVar string = "Genesis"
	fmt.Println(constVar)

	const piConst float32 = 3.14615
	fmt.Println(piConst)
}
