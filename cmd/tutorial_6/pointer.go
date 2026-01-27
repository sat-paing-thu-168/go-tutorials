package main

import (
	"fmt"
)

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("Memory Location of thing2 array is: %p \n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p point is %v \n", *p)
	fmt.Printf("The value if i is %v \n", i)
	p = &i
	*p = 1
	fmt.Printf("The value p point is %v \n", *p)
	fmt.Printf("The value if i is %v \n", i)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 5
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	var thing = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Memory Location of thing1 array is: %p \n", &thing)
	fmt.Println(square(&thing))
	fmt.Println(thing)

	k, j := 4, 49
	fmt.Println(&k, &j)

}
