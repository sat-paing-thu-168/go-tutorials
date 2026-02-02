package main

// itearting channel only stop when channel is closed

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"Apple", "Banana", "Papaya", "Orange", "Grape"}
var result = make([]string, 0, 5)

func main() {
	wg := sync.WaitGroup{}

	channel := make(chan string, len(dbData))

	t0 := time.Now()
	for _, fruit := range dbData {
		wg.Add(1)
		go dbCall(fruit, &wg, channel)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
	for fruit := range channel {
		result = append(result, fruit)
		log(fruit)
	}

	fmt.Printf("Total Execution time is %v\n", time.Since(t0))
}

func dbCall(fruit string, wg *sync.WaitGroup, channel chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	channel <- fruit

}

func log(fruit string) {
	fmt.Println("Successfully added fruit: ", fruit)
}
