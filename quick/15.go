package main

import (
    "fmt"
    "net/url"
)

type Node struct {
    name string
}

func test(c interface{}) {
    obj := c.(Node)
    obj.name="ok"
    fmt.Printf("testtest=%+v\r\n", obj)

}



func main() {
    xx := Node{
        name: "hzc",
    }
    test(xx)


    providerID := "openstack:///d2e9a984-8faf-4dba-a770-63d2015e8d33"
    parsedID, err := url.Parse(providerID)
    if err != nil {
        fmt.Print(err)
        return
    }

    if parsedID.Scheme != "openstack" {
        fmt.Printf("unrecognized provider %q", parsedID.Scheme)
        return
    }

    fmt.Printf("the scheme=%q id=%q\r\n", parsedID.Scheme, parsedID.Host)


//    loop_i := [5]string {"i_hu", "i_zheng", "i_chuan", "i_hello", "i_world"}
//    loop_j := [5]string {"j_hu", "j_zheng", "j_chuan", "j_hello", "j_world"}

    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            if j == 3 {
                fmt.Println("break out")
                break
            }
            fmt.Printf("\t j=%d\r\n", j)
        }
        fmt.Printf("i=%d\r\n", i)
    
    }

}
