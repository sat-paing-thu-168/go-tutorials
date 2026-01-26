/*
Exercise 2: The "Slice Surgeon"

The Goal: Master the [low:high] syntax.

Given this starting slice: planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

Write a program to:

Create a new slice rockyPlanets by slicing the planets slice (should contain Mercury through Mars).

Create a new slice gasGiants by slicing the planets slice (should contain Jupiter through Neptune).

The Twist: Change "Mars" to "REDACTED" inside the rockyPlanets slice.

Print the original planets slice. Does it show "Mars" or "REDACTED"? (This tests if you understand that slices share the same underlying array).
*/

package main

import "fmt"

func main() {
	intSlice := []int{}
	for index := range 20 {
		intSlice = append(intSlice, index+1)
		fmt.Printf("Len :%d, Cap: %d, Address: %p \n", len(intSlice), cap(intSlice), &intSlice[0])
	}
	fmt.Println(intSlice)
}
