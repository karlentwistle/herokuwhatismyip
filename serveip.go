package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
    http.HandleFunc("/", serveIP)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func serveIP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, r.Header.Get("X-Forwarded-For"))
}
