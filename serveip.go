package main

import (
  "fmt"
  "net/http"
  "strings"
  "os"
)

func main() {
    http.HandleFunc("/", serveIP)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func serveIP(w http.ResponseWriter, r *http.Request) {
    ip := strings.Split(r.RemoteAddr, ":")
    fmt.Fprintf(w, ip[0])
}
