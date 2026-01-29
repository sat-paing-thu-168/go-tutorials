/*
Exercise 1: The "Delayed Greeting" (Basics)

Goal: Prove that Goroutines run independently of the main function.

    Write a function sayHello(name string) that sleeps for 1 second and then prints "Hello, [name]".

    In main, call this function as a Goroutine using the go keyword for three different names.

    The Challenge: If you run the program now, it will likely exit before anything prints.
	 Use time.Sleep in main for 2 seconds to keep the program alive long enough to see the results.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go sayHello("Sat Paing")
	wg.Wait()
	fmt.Println("Hello Greetings")
}

func sayHello(name string) {
	var delay float32 = rand.Float32() * 3000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Hello ", name)
	wg.Done()
}
