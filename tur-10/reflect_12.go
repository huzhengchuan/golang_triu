package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Name string 
    code int
}


func main() {
    p := new(User)
    v := reflect.ValueOf(p).Elem()

    name := v.FieldByName("Name")
    code := v.FieldByName("code")

    fmt.Printf("name: canadr = %v, canset=%v\n", name.CanAddr(), name.CanSet())
    fmt.Printf("code: canadr = %v, canset=%v\n", code.CanAddr(), code.CanSet())

    if name.CanSet() {
        name.SetString("Tom")
    }
}

