package main

import (
    "fmt"
    "reflect"
)

type  X int

func (X) String() string {
    return ""
}


func main() {

    var a X
    t := reflect.TypeOf(a)

    st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
    fmt.Println(t.Implements(st))

    it := reflect.TypeOf(0)
    fmt.Println(t.ConvertibleTo(it))

    fmt.Println(t.AssignableTo(st), t.AssignableTo(it))
}

