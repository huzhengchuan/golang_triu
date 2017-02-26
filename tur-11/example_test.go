package main

import (
    "fmt"
)


func add(x, y int) int {
    return x + y
}

func ExampleAdd() {
    fmt.Println(add(1, 2))
    fmt.Println(add(2,2))

    // Output:
    // 3
    // 4
}
