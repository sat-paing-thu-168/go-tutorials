package main

import (
	"errors"
	"fmt"
)

func main() {
	name := "Sae Paing"
	printMe(name)

	numerator := 10
	denominator := 5
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// } else {
	// 	fmt.Printf("Result of the integer is %v with remainder %v", result, remainder)
	// }
	switch {
	case err != nil:
		fmt.Print(err.Error())
	default:
		fmt.Printf("Result of the integer is %v with remainder %v", result, remainder)
		fmt.Println()
	}
	switch remainder {
	case 0:
		printMe("Yes Remainder is Zero")
	}
}

func printMe(name string) {
	fmt.Println(name)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide By Zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
