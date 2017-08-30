package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("hello world")

	per := person(name: "hello", age: 24)
	fmt.Println(per)
}
