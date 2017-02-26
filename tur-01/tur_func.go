package main

import (
    "errors"
    "fmt"
)

func div(a, b int) (int , error) {
    if b == 0 {
        return 0, errors.New("division bay zero")
    }

    return a/b, nil
}


func test(x int) func() {
    return func() {
        println(x)
    }
}

func deployFun(a, b int) {
    defer println("dispose.......")

    println(a/b)
}

func main() {
    fmt.Println("############testcase-01")
    a, b := 10, 0
    c, err := div(a, b)

    fmt.Println(c, err)
    
    fmt.Println("############testcase-02")
    x := 100
    f := test(x)
    f()

    fmt.Println("############testcase-03")
    deployFun(10, 0)
}
