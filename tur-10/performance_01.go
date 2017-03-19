package main

import (
    "reflect"    
)

type Data struct {
    X int
}

var d = new(Data)

func  set(x int) {
    d.X = x
}

func reset(x int) {
    v := reflect.ValueOf(d).Elem()
    f := v.FieldByName("X")
    f.Set(reflect.ValueOf(x))
}

func BenchmardkSet(b *testing.B) {
    for i := 0; i < b.N; i++ {
        set(100)
    }
}

func BenchmardkRset(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rset(100)
    }
}

