package main

import (
    "fmt"
)

func main() {

    fmt.Println("Hello world!")
    var str []string = make([]string, 10)
    str[0] = "*"
    str[1] = "api"
    fmt.Printf("%v\r\n", str)

    GetValidSource(str)

}

func GetValidSource(sources []string) ([]string, error) {
    
    validated := make([]string, 0, 10)
    for _, source := range sources {
        switch source {
        
            case "*", "api":
                fmt.Println("can fly")
                validated = append(validated, source)
                break
            case "":
                break
        }
    }

    fmt.Printf("hzc:%v", validated)
    return validated, nil

}
