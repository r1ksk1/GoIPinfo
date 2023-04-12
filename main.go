package main

import (
    "fmt"
    "net"
    "net/http"
)

const (
    dnsServer1 = "1.1.1.1"
    dnsServer2 = "1.0.0.1"
    dnsServer1 = "8.8.8.8"
    dnsServer2 = "8.8.4.4"
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

    http.ListenAndServe(":8080", nil)
}

func lookupHostname(ipAddr string) (string, error) {
    servers := []string{dnsServer1, dnsServer2, dnsServer3, dnsServer4}
    for _, server := range servers {
        config := &net.ResolverConfig{Nameservers: []string{server}}
        resolver := &net.Resolver{Config: config}

        addr, err := net.ResolveIPAddr("ip", ipAddr)
        if err != nil {
            return "", err
        }

        names, err := resolver.LookupAddr(addr.String())
        if err == nil && len(names) > 0 {
            return names[0], nil
        }
    }

    return "", fmt.Errorf("could not resolve hostname for IP address %s", ipAddr)
}
