package main


import (
    "fmt"
)

type Base struct {
    name string
}

type Node struct {
    basees []*Base
    cc  int
}

func main() {
    node1 := new(Node)
    fmt.Printf("%v\r\n", node1.basees)

}
