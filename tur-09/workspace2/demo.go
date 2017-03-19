package main

import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte(`hello world`))
            })
    http.ListenAndServe(":3000", nil) // <-今天讲的就是这个ListenAndServe是如何工作的
}
