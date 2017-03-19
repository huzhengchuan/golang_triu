package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,` 
        <meta name="go-import" content="qyuhen.com/test git https://github.com/qyuhen/test">
        `)
}

func main() {
    http.HandleFunc("/test", handler)
    http.ListenAndServe(":80", nil)
}
