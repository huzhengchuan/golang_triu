package main

import (
    "fmt"
)

type Node struct {
    name string
}

func main() {
    nodes := make([]Node, 1,1)
    nodes[0] = Node{name:"test"}
    
    index:=0
    nodes[index] = nodes[len(nodes) - 1]
    nodes = nodes[:len(nodes)-1]
    fmt.Printf("%+v  %d\r\n", nodes, len(nodes))
}
