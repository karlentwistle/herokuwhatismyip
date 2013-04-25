package main

import (
  "fmt"
  "net/http"
  "os"
  "strings"
)

func main() {
    http.HandleFunc("/", serveIP)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func serveIP(w http.ResponseWriter, r *http.Request) {
  if len(r.Header.Get("X-Forwarded-For")) > 0 {
    ips := strings.Split(r.Header.Get("X-Forwarded-For"), ", ")
    fmt.Fprintf(w, ips[(len(ips) - 1)])
  } else {
    ips := strings.Split(r.RemoteAddr, ":")
    fmt.Fprintf(w, ips[0])
  }

}
