package main

import (
	"fmt"
	"time"
)

func main() {
	var intArr [3]int32

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])

	intArr2 := [...]int32{1, 2, 3}
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println()
	fmt.Println("=====")

	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println()
	fmt.Println("=====")

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println("=====")

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
	fmt.Println("=====")

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	fmt.Println("=====")

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	// delete(myMap2, "Adam")
	fmt.Println(myMap2)
	var age, ok = myMap2["Sarah"]
	if ok {
		fmt.Printf("The age is %v", age)
		fmt.Println()
	} else {
		fmt.Println("Invalid Name")
	}
	fmt.Println("=====")

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age %v", name, age)
		fmt.Println()
	}
	fmt.Println("=====")

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
