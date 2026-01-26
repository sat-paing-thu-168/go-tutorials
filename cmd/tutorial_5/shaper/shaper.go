/*
	Try to create an interface called Shaper with an Area() float64 method.
	Then create two structs, Circle and Square, and make them "implement" that interface.
*/

package main

import "fmt"

type Shaper interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	side1 float64
	side2 float64
}

func (c Circle) Area() float64 {
	pi := 3.142
	return c.radius * c.radius * pi
}

func (s Square) Area() float64 {
	return s.side1 * s.side2
}

func calculateArea(s Shaper) {
	fmt.Printf("The area for the Shpae is %v\n", s.Area())
}

func describeShaper(s Shaper) {
	switch v := s.(type) {
	case Circle:
		fmt.Printf("This is a cricle with radius %v \n", v.radius)
	case Square:
		fmt.Printf("This is a square with sides %v and %v\n", v.side1, v.side2)
	default:
		fmt.Println("Unkown Shape")
	}

}

func main() {
	var circle1 Circle = Circle{4}
	var circle2 Circle = Circle{5}
	var square Square = Square{2, 3}
	// describeShaper(square)
	// calculateArea(square)
	// describeShaper(circle1)
	// calculateArea(circle2)

	shapers := make([]Shaper, 0, 3)
	appendShapers := []Shaper{circle1, circle2, square}
	shapers = append(shapers, appendShapers...)

	for _, shape := range shapers {
		calculateArea(shape)
	}
}
