package main

import (
    "fmt"
    "reflect"
)
type x int

func main() {

    var a x = 100
    t := reflect.TypeOf(a)

    fmt.Println(t.Name(), t.Kind())
}
