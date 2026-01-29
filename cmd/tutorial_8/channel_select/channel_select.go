package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	time.Sleep(time.Second * 3)
}

func test1(c chan<- string) {
	for {
		pause()
		c <- "hello"
	}
}

func test2(c chan<- string) {
	for {
		pause()
		c <- "world"
	}
}

func main() {
	rand.NewSource(time.Now().Unix())

	c1 := make(chan string)
	c2 := make(chan string)

	go test1(c1)
	go test1(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
