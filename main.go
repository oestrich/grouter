package main

import (
  "fmt"
  "http"
  "flag"
  "router"
  "os"
)

var port = flag.Int("p", 3000, "port to host on")
var test = flag.Bool("t", false, "run tests")

func main() {
  flag.Parse()

  if *test {
    router.TestParsing()
    router.TestRouterParsing()
    os.Exit(0)
  }

  router.WebRouter.Append(router.Route{Path: "/posts/:id"})

  router.SetupRouter()

  addr := fmt.Sprint(":", *port)

  fmt.Println("Web server started")
  if err := http.ListenAndServe(addr, nil); err != nil {
    fmt.Println("Couldn't launch server: ", err)
  }
}
