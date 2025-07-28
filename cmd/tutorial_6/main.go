package main

import "fmt"

// Structs are a way of defining your own type
type menuItem struct {
	title     string
	price     int32
	options   []string
	itemOwner owner // or just "owner"
	owner2          // direct access to fields of owner2
	string
	quantity int
}

type owner struct {
	name string
}

type owner2 struct {
	name string
}

// Function for struct, it's like calling a class's method
func (e menuItem) getTotalPrice() int {
	return e.quantity * int(e.price)
}

func isExpensive(e menuItem, o owner) bool {
	if o.name == "A" && e.price >= 0 {
		return false
	}

	return true
}

func main() {
	var myMenuItem menuItem
	fmt.Println(myMenuItem.title, myMenuItem.price, myMenuItem.options)

	var myMenuItem2 menuItem = menuItem{title: "Kebap", price: 50, options: []string{"A", "B"}}
	fmt.Println(myMenuItem2.title, myMenuItem2.price, myMenuItem2.options)

	// You can omit the field names, and assigned them in order
	var myMenuItem3 menuItem = menuItem{
		"Kebap",
		50,
		[]string{"A", "B"},
		owner{"Genesis"},
		owner2{"OWNER2"},
		"STRING",
		10,
	}
	fmt.Println(
		myMenuItem3.title,
		myMenuItem3.price,
		myMenuItem3.options,
		myMenuItem3.itemOwner.name,
		myMenuItem3.name,        // Shorter Alternative
		myMenuItem3.owner2.name, // or use it like this
		myMenuItem3.string,
		myMenuItem3.quantity,
	)

	var myMenuItem4 menuItem
	myMenuItem4.price = 1453
	myMenuItem4.title = "Taco"
	myMenuItem4.options = []string{"Hot", "None", "Saury"}
	fmt.Println(myMenuItem4.title, myMenuItem4.price, myMenuItem4.options)

	fmt.Println("myMenuItem3.getTotalPrice(): ", myMenuItem3.getTotalPrice())
	fmt.Println("isExpensive: ", isExpensive(myMenuItem, myMenuItem.itemOwner))

	var myGasEngine gasEngine = gasEngine{2, 1}
	var myElectricEngine electricEngine = electricEngine{12, 35}
	canMakeIt(myGasEngine, 50)
	canMakeIt(myElectricEngine, 100)
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkw uint8
	kwh  uint8
}

// Interfaces are generic
// Here, engine can be electric or gas
type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You can NOT make it")
	}
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkw
}
