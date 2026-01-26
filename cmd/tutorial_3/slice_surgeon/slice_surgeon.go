/*
Exercise 3: The Filter Shortcut

The Goal: Use range and append effectively.

Create a slice with the following numbers: [1, 5, 2, 7, 4, 10, 15, 6].

Write a loop that iterates through this slice and only appends the even numbers to a new slice called evens.

Print the evens slice.

Bonus: Try to do this using only one for range loop.
*/

package main

import "fmt"

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	rockyPlantes := planets[0:4]
	gasGiants := planets[4:8]
	rockyPlantes[3] = "REDACTED"
	fmt.Println(planets)
	fmt.Println(rockyPlantes)
	fmt.Println(gasGiants)
}
