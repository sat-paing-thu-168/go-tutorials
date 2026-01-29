package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string, 2)

	go func() {
		time.Sleep(time.Second + 1)
		msgChannel <- "Hello"
		fmt.Println("Printing 1st")
		msgChannel <- "World"
		fmt.Println("Printing Second")
	}()

	msg1 := <-msgChannel
	fmt.Println("Adding 1st")
	msg2 := <-msgChannel
	fmt.Println("Adding 2nd")

	fmt.Println(msg1, msg2)
}
