package main

import (
    "fmt"
    "net"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        ipAddr := r.RemoteAddr
        hostname, _ := lookupHostname(ipAddr)
        userAgent := r.Header.Get("User-Agent")
        language := r.Header.Get("Accept-Language")

        fmt.Fprintf(w, "Client IP address: %s\n", ipAddr)
        fmt.Fprintf(w, "Client hostname: %s\n", hostname)
        fmt.Fprintf(w, "Client user agent: %s\n", userAgent)
        fmt.Fprintf(w, "Client language: %s\n", language)
    })

    http.ListenAndServe(":8448", nil)
}

func lookupHostname(ipAddr string) (string, error) {
    addr, err := net.ResolveIPAddr("ip", ipAddr)
    if err != nil {
        return "", err
    }

    names, err := net.LookupAddr(addr.String())
    if err != nil || len(names) == 0 {
        return "", err
    }

    return names[0], nil
}
