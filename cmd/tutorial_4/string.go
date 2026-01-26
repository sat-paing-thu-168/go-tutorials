package main

import "fmt"

func main() {
	myString := []rune("resume")
	indexed := myString[0]
	fmt.Printf("%v, %t\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Println(len(myString))

}
