package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    log.Println("Listening on :443 and :80...")
    go http.ListenAndServe(":80", nil)
    err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/thomashansen.xyz/fullchain.pem", "/etc/letsencrypt/live/thomashansen.xyz/privkey.pem", nil)
    if err != nil {
        log.Fatal(err)
    }
}

