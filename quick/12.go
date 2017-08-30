package main

import (
    "fmt"
)


func test(a string, b string) {
    fmt.Printf("print %q\r\n", a)
}

func main() {
    fmt.Println("testing.....")
    test("huzhengchuan", "bb")
}
