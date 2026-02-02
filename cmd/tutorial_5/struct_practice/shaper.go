/*
	Try to create an interface called Shaper with an Area() float64 method.
	Then create two structs, Circle and Square, and make them "implement" that interface.
*/

package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float64
}

func Area(s Shaper) float64 {
	return s.Area()
}

type Circle struct {
	radius float64
}

type Square struct {
	side1 float64
	side2 float64
}

func (c Circle) Area() float64 {
	PI := 3.142
	area := math.Pow(c.radius, 2) * PI
	return area
}

func (s Square) Area() float64 {
	return s.side1 * s.side2
}

func main() {
	circle := Circle{radius: 2}
	square1 := Square{side1: 4, side2: 3}
	square2 := Square{side1: 5, side2: 5}

	shapes := []Shaper{circle, square1, square2}
	calculateArea(shapes)
}

func calculateArea(shapes []Shaper) {
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Radius is %v, Area of Circle is %v \n", s.radius, Area(shape))
		case Square:
			fmt.Printf("Side1 length is %v, Side2 length is %v, Area of Circle is %v \n", s.side1, s.side2, Area(shape))
		}
	}

}
