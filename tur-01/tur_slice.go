package main

import (
    "fmt"
)


func funSlice() {

    x := make([]int , 0, 5)

    for i := 0; i < 8; i++ {
        x = append(x, i)
    }

    fmt.Println(x)

}

func funMap() {
    m := make(map[string]int)

    m["a"] = 1
    x, ok := m["b"]
    fmt.Println(x, ok)

    delete(m, "a")
}


func main() {
    
    fmt.Println("####### test-case -0")
    funSlice()
    fmt.Println("####### test-case -1")
    funMap()
}
