package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	aph   uint8
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.aph
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first")
	}
}

func main() {

	var myEnginer2 = struct {
		mpg     uint8
		gallons uint8
	}{40, 50}
	fmt.Println(myEnginer2.mpg, myEnginer2.gallons)

	var customEngine gasEngine = gasEngine{50, 50}
	fmt.Printf("Total miles left in tank: %v\n", customEngine.milesLeft())

	var defaultEngine electricEngine = electricEngine{25, 15}
	canMakeIt(defaultEngine, 50)
}
