package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6"}
var results = make([]string, 0, 6)

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total Execution time is : %v", time.Since(t0))
	fmt.Printf("\n Results are : %v", results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000

	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("The result from db is: ", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\n Current results are : %v", results)
	m.RUnlock()
}
