package main

import (
)

func mkslice() []int {
    s := make([]int, 0, 10)
    s = append(s, 100)
    return s
}

func mkmap() map[string] int {
    m := make(map[string] int)
    m["a"] = 1
    return m
}

func mknewError() {
    p := new(map[string] int)
    m := *p
    m["a"] = 1
    println(m)
}

func main() {
    m := mkmap()
    println(m["a"])

    s := mkslice()
    println(s[0])

    mknewError()
}
