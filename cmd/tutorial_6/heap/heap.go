package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() *person {
	m := person{name: "Sat Paing", age: 26}
	fmt.Printf("Init Person ---> %p\n", &m)
	return &m
}

func main() {
	fmt.Printf("Main ----> %p\n", initPerson())

	k, j := 4, 49
	fmt.Println(&k, &j)
	c := &k
	fmt.Println(*c)
	fmt.Printf("%T\n", c)
	*c = 223
	fmt.Println(k)

	c = &j
	*c = *c / 37
	fmt.Println(j)
}
