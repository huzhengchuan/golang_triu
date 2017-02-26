package main

import (
    "testing"
    "time"
    "os"
)


func add(x, y int) int {
    return x + y
}


func TestA(t *testing.T) {
    t.Parallel()
    time.Sleep(time.Second*2)
}

func TestB(t *testing.T) {
    if os.Args[len(os.Args) - 1] == "b" {
        t.Parallel()
    }

    time.Sleep(time.Second *2)
}

func TestAdd(t *testing.T) {
    if add(1, 2) != 3 {
        t.FailNow()
    }
}
