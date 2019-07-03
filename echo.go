package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, r.RemoteAddr);
}

func main() {
    http.HandleFunc("/", handler)

    port := ":26256"
    log.Printf("Serving on %s...\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}
