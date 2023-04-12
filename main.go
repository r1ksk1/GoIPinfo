package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "IP address: %s\n", r.RemoteAddr)
        fmt.Fprintf(w, "Hostname: %s\n", r.Host)
        fmt.Fprintf(w, "User-Agent: %s\n", r.Header.Get("User-Agent"))
        fmt.Fprintf(w, "Language: %s\n", r.Header.Get("Accept-Language"))
    })

    http.ListenAndServe(":8448", nil)
}
