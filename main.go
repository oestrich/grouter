package main

import (
  "fmt"
  "http"
  "flag"
  "router"
)

var port = flag.Int("p", 3000, "port to host on")

func main() {
  flag.Parse()

  http.HandleFunc("/", router.Router)
  addr := fmt.Sprint(":", *port)

  fmt.Println("Web server started")
  if err := http.ListenAndServe(addr, nil); err != nil {
    fmt.Println("Couldn't launch server: ", err)
  }
}
