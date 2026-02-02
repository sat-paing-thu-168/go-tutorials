package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeStore struct {
	mutex  sync.Mutex
	result []string
}

func (store *SafeStore) save(fruit string) {
	store.mutex.Lock()
	store.result = append(store.result, fruit)
	store.mutex.Unlock()
}

var dbData = []string{"Apple", "Banana", "Papaya", "Orange", "Grape"}

func main() {
	wg := sync.WaitGroup{}
	store := SafeStore{result: make([]string, 0, 5)}

	t0 := time.Now()
	for _, fruit := range dbData {
		wg.Add(1)
		go dbCall(fruit, &wg, &store)
	}
	wg.Wait()
	fmt.Printf("Total Execution time is %v\n", time.Since(t0))
}

func dbCall(fruit string, wg *sync.WaitGroup, store *SafeStore) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	store.save(fruit)
	log(fruit)
}

// func save(fruit string, mutex *sync.Mutex) {
// 	mutex.Lock()
// 	result = append(result, fruit)
// 	mutex.Unlock()
// }

func log(fruit string) {
	fmt.Println("Successfully added fruit: ", fruit)
}
