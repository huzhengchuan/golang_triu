package main

import (
    "sets"
)

var disablecontroller = map[string]Empty(
    "bootstrapsigner",
    "tokencleaner",
)

func main() {
    fmt.Println(disablecontroller.List())
}
