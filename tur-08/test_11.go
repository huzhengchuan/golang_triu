package main

import (
    
)

func main() {
    c := make(chan int, 3)

    c <- 1
    c <- 2
    c <- 3
    c <- 4
    println(<-c)
    println(<-c)
    
}
