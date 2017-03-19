package main

import (
    
)

type Data struct {
    X int
}

func (x *Data) Inc() {
    x.X++
}

var d = new(Data)
var v = reflect.ValueOf(d)
var n = v.MethodByName("Inc")


func call() {
    d.Inc()
}

func rcall() {
    m.Call(nil)
}

func main() {
    
}
