package main


import (
    "fmt"
)

type Test interface {
    printok(str string) error
}

type People struct {
}

func (hzc *People)printok(str string) error {
    fmt.Printf("name is %s", str)
    return nil
}

func getObj() Test {
    return nil
}
func main() {

    fmt.Println("hello world")
    peopleOne := getObj()
    peopleOne.printok("nihao, huzhengchuan")
}
