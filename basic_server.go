// A very simple, basic file server which can be used in 
// testing  that doesn't require TLS setup or anything
// localhost:8080
package main

import (
    "log"
    "net/http"
    "fmt"
)


func main() {
    var port int = 8080
    var port_s string = fmt.Sprintf(":%d", port)

    fmt.Printf("Starting server at port %d\n", port)

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    if err := http.ListenAndServe(port_s, nil); err != nil {
        log.Fatal(err)
    }
}

