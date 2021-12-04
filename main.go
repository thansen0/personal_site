package main

import (
    "log"
    "net/http"
    "os"
)

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://thomashansen.xyz:443"+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    log.Println("Listening on :443 and redirecting :80...")

    go func(){
        errhttps := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/thomashansen.xyz/fullchain.pem", "/etc/letsencrypt/live/thomashansen.xyz/privkey.pem", nil)

        if errhttps != nil {
            log.Fatal(errhttps)
            os.Exit(3)
        }
    }()

    err := http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))
    if err != nil {
        log.Fatal(err)
        os.Exit(3)
    }
}

