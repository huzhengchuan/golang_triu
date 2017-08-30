package main

import (
	"fmt"
)

type base struct{
	name string
	age int
}

type student struct{
	x base
}

func main() {
	fmt.Println("Hello golang!")
	var stu *student = new(student)
	stu.x.name = "hzc"
	stu.x.age = 15
	fmt.Print(stu)
}
