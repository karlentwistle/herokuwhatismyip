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
    //ip := strings.Split(r.RemoteAddr, ":")
    fmt.Println(r.Header)
    fmt.Fprintf(w, "hey")
}
