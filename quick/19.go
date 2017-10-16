package main


import (
    "fmt"
)

type clientset struct {
    name *string
}

func (c *clientset) corev1() {
    fmt.Print("hello\n")
}

func getObj() *clientset{
    return nil
}


func main() {
    p := getObj()
    p.corev1()    
    fmt.Print(p.name)
}
