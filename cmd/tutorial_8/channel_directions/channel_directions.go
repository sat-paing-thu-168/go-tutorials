package main

import (
	"fmt"
	"time"
)

func writer(channel chan<- string, msg string) {
	channel <- msg

}

func reader(channel <-chan string) {
	msg := <-channel
	fmt.Println(msg)
}

func main() {
	messageChannel := make(chan string, 1)

	go reader(messageChannel)

	for i := 0; i < 10; i++ {
		go writer(messageChannel, fmt.Sprintf("msg %d", i))
	}
	time.Sleep(time.Second + 1)
}
