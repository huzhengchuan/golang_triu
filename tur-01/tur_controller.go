package main

//import (
//    "fmt"
//)

func funIf() {
    x := 100

    if x > 0 {
        println("x")
    } else if x < 0 {
        println("-x")
    } else {
        println("0")
    }
}

func funSwitch() {
    
    x := 100

    switch {
        case x > 0:
            println("x")
        case x < 0:
            println("-x")
        default:
            println("0")
    }
}

func funFor1() {
    for i := 0; i < 5; i++ {
        println(i)
    }

    for i := 4; i >= 0; i-- {
        println(i)
    }
}

func funFor2() {
    x := 0

    for x < 5 {
        println(x)
        x++
    }
}

func funFor3() {
    x := 4

    for {
        println(x)
        x--

        if x < 0 {
            break
        }
    }
} 

func funRange() {
    x := []int{100, 101, 102}

    for i, n := range x {
        println(i, ":", n)
    }
}

func main() {
   funIf()
   funSwitch()
   funFor1()
   funFor2()
   funFor3()
}
