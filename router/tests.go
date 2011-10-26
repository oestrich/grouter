package router

import (
  "fmt"
)

func TestParsing() {
  route := &Route{Path: "/posts/:id"}
  pass, handler, vars := route.ParseRoute("/posts/1")
  fmt.Println(handler)
  fmt.Println(vars)

  if vars[":id"] != "1" {
    fmt.Println("ERROR")
  }

  pass, _, _ = route.ParseRoute("/posts/another/path")
  if pass != false {
    fmt.Println("ERROR")
  }

  pass, _, _ = route.ParseRoute("/comments/1")
  if pass != false {
    fmt.Println("ERROR")
  }
}
