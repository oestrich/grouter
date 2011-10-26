package router

import (
  "fmt"
)

func TestParsing() {
  route := &Route{Path: "/posts/:id/:format"}
  handler, vars := route.ParseRoute("/posts/1")
  fmt.Println(handler)
  fmt.Println(vars)

  if vars[":id"] != "1" {
    fmt.Println("ERROR")
  }
}
