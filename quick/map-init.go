package main

import (
    "fmt"
)

func main() {
    fmt.Println("start working")

    var m1 map[string]string
    m1 = make(map[string]string)
    m1["a"] = "aa"
    m1["b"] = "bb"
    fmt.Println("test ", m1)

    m2 := make(map[string]string)
    m2["a"] = "aa"
    m2["b"] = "bb"
    fmt.Println("test ", m2)


    m3 := map[string]string {
        "a" : "aa",
        "b" : "bb",
    }
    fmt.Println("test ", m3)

    if v, ok := m1["a"]; ok {
        fmt.Println("get result ", v)
    } else {
        fmt.Println("key not found")
    }

    for k, v := range m1 {
        fmt.Println(k, v)
    }

}
