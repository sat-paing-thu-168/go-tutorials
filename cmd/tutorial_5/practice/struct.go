package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

type electricEngine struct {
	mpkwh uint8
	aph   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.aph
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You better fuel up")
	}
}

type owner struct {
	name string
}

func main() {
	var myEngine gasEngine = gasEngine{10, 20}
	fmt.Println(myEngine.milesLeft())
	fmt.Printf("MPG %v , Gallons %v\n", myEngine.mpg, myEngine.gallons)

	var myEngine2 electricEngine = electricEngine{30, 40}
	fmt.Println(myEngine2.milesLeft())
	fmt.Printf("MPG %v , Gallons %v\n", myEngine2.mpkwh, myEngine2.aph)

	canMakeIt(myEngine, 210)
	canMakeIt(myEngine2, 30)
}
