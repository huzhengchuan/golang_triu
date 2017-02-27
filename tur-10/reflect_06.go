package main

import (
    "fmt"
    "reflect"
)


type user struct {
    name string
    age int
}

type manager struct {
    user
    title string
}

func main() {
    
    var m manager
    
    t := reflect.TypeOf(m)

    name, _ := t.FieldByName("name")
    fmt.Println(name.Name, name.Type)

    age := t.FieldByIndex([]int {0, 1})
    fmt.Println(age.Name, age.Type)
}
