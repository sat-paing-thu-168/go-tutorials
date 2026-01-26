/*
Challenge accepted! Here is a three-part "Go Fundamentals" homework set designed to test your understanding of slices, loops, and that tricky capacity growth we just talked about.
Exercise 1: The "Growth Tracker"

The Goal: See the memory reallocations with your own eyes.

Create an empty slice of integers.

Use a for loop to append the numbers 1 through 20 to that slice one by one.

Inside the loop, print the length, capacity, and the memory address of the first element of the slice.

Hint: Use fmt.Printf("Len: %d, Cap: %d, Address: %p\n", len(s), cap(s), &s[0])

Observation: Notice at which specific numbers the address changes. Every time the address changes, Go has "moved house" to a bigger memory block.
*/

package main

import "fmt"

func main() {
	slice := []int{1, 5, 2, 7, 4, 10, 15, 6}
	evens := make([]int, 0, 5)
	for _, value := range slice {
		if value%2 == 0 {
			evens = append(evens, value)
		}
	}
	fmt.Println(evens)
}
