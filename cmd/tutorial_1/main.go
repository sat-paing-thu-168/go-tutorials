package main

import (
	"fmt"
	"unicode/utf8"

	"rsc.io/quote"
)

func main() {
	var intNum int8 = 1
	fmt.Println(intNum)
	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var num1 float32 = 10.1
	var num2 int32 = 2
	var result float32 = num1 + float32(num2)
	fmt.Println(result)

	const myText string = "H"
	fmt.Println(myText)
	fmt.Println(utf8.RuneCountInString(myText))

	var status uint8
	fmt.Println(status)

	randomText1, randomText2 := "lorem", "Ipsum"
	fmt.Println(randomText1)
	fmt.Println(randomText2)

	fmt.Print(quote.Go())
}
