package main

import (
	"fmt"
	"sync"
	"time"
)

func pause() {
	time.Sleep(time.Second + 1)
}

func sendMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	pause()
	fmt.Println(msg)
}

func main() {
	// msgChan := make(chan string)

	var wg = sync.WaitGroup{}

	wg.Add(3)

	go func(msg string) {
		defer wg.Done()
		pause()
		// msgChan <- "Channel Input 1"
		// msgChan <- "Channel Input 2"
		fmt.Println(msg)
	}("New Test")

	go sendMessage("hello", &wg)

	go sendMessage("test1", &wg)

	wg.Wait()
}
