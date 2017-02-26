package main

import (
    "fmt"
    "math"
    "strconv"
)
const y = 0x200

func main() {
    fmt.Println("test-case--01")
    println(y)
    

    fmt.Println("test-case--02")
    a, b, c := 100, 0144, 0x64
    fmt.Println(a, b, c)
    fmt.Printf("0b%b, %#o, %#x\n", a, a, a)
    fmt.Println(math.MinInt8, math.MaxInt8)

    fmt.Println("test-case--03")
    a1, _ := strconv.ParseInt("1100100", 2, 32)
    b1, _ := strconv.ParseInt("0144", 8, 32)
    c1, _ := strconv.ParseInt("64", 16, 32)
    fmt.Println(a1, b1, c1)
    
    println("0b" + strconv.FormatInt(a1, 2))
    println("0" + strconv.FormatInt(a1, 8))
    println("0x" + strconv.FormatInt(a1, 16))

    fmt.Println("test-case--03")
    var a2 byte = 0x11
    var b2 uint8 = a2
    var c2 uint8 = a2 + b2
    println(c2)

}
