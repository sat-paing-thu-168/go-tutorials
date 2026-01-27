package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbItems = []string{"Apple", "Banana", "Dragonfruit", "Elderberry", "Fig", "Guava"}
var results = make([]string, 0, 5)

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbItems); i++ {
		wg.Add(1)
		go fetchFruits(i)
	}
	wg.Wait()
	fmt.Printf("\n Total Execution time is : %v", time.Since(t0))
	fmt.Printf("\n Results are : %v", results)
}

func fetchFruits(index int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbItems[index])
	log()
	wg.Done()
}

func save(fruit string) {
	m.Lock()
	results = append(results, fruit)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\n Current results are : %v", results)
	m.RUnlock()
}
