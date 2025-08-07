package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// If we write like the following,
// once we do:
// 
// data, _ := json.Marshal(ci)
// fmt.Println(string(data)) 
// 
// This is gonna be like:
// {"name":"Süha","email":"suha@example.com"}
// Keys will start with small letter
// 
// type contactInfo struct {
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

type contactInfo struct {
	Name  string
	Email string
}
// {"Name":"Süha","Email":"suha@example.com"}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func jsonExample() {
	// Here we need to specify the type of the struct we want to load
	// Because the compiler cannot infer the type of the struct
	// So we write "[contactInfo]"
	var contacts []contactInfo = loadJSON[contactInfo]("./contacts.json")
	fmt.Println("contacts: ", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchases.json")
	fmt.Println("purchases: ", purchases)

}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := os.ReadFile(filePath)

	var loaded = []T{}
	// Unmarshaling means parsing/loading JSON data into a struct
	json.Unmarshal(data, &loaded)

	return loaded
}
