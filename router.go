package main

import (
  "fmt"
  "io"
  "http"
  "flag"
)

func Router(writer http.ResponseWriter, r *http.Request) {
  io.WriteString(writer, "Hit a page\n")
  io.WriteString(writer, r.Method + "\n")
  io.WriteString(writer, r.URL.Raw + "\n")
}

var port = flag.Int("p", 3000, "port to host on")

func main() {
  flag.Parse()

  http.HandleFunc("/", Router)
  addr := fmt.Sprint(":", *port)

  fmt.Println("Web server started")
  if err := http.ListenAndServe(addr, nil); err != nil {
    fmt.Println("Couldn't launch server: ", err)
  }
}
