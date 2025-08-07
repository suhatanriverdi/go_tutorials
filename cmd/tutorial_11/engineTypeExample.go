package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func engineExample() {
	var gasCar = car[gasEngine]{
		carMake:  "Toyota",
		carModel: "Corolla",
		engine: gasEngine{
			gallons: 10,
			mpg:     30,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model S",
		engine: electricEngine{
			kwh:   100,
			mpkwh: 300,
		},
	}

	fmt.Println("Gas", gasCar)
	fmt.Println("Electric", electricCar)
}
