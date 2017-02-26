package main

import (
    "time"
    "testing"
)

func BenchmarkAdd(b *testing.B) {
    
    time.Sleep(time.Second)
    b.ResetTimer()


    for i := 0; i < 100; i++ {
        _ = add(1, 2)

        if i == 1 {
            b.StopTimer()
            time.Sleep(time.Second)
            b.StartTimer()
        }
    }
}
