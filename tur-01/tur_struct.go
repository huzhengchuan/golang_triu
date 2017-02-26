package main

import (
    "fmt"
)

type user struct {
    name string
    age byte
}

type manager struct {
    user
    title string
}


type X int
func (x *X) inc() {
    *x++
}



func (u user) ToString() string {
    return fmt.Sprintf("%+v", u)
}



func main() {

    fmt.Println("testcase ----2")
    var m manager

    m.name = "Tom"
    m.age = 29
    m.title = "CTO"

    fmt.Println(m)

    fmt.Println("testcase ----2")
    var x X
    x.inc()
    println(x)


    fmt.Println("testcase -----3")
    var man manager
    man.name = "Hzc"
    m.age = 20
    fmt.Println(m.ToString())
}
