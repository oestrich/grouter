package router

import (
  "fmt"
)

func TestParsing() {
  route := &Route{Path: "/posts/:id"}
  pass, _, vars := route.ParseRoute("/posts/1")

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

func TestRouterParsing() {
  fmt.Println("\nStarting Router Tests\n")

  router := &Router{}
  router.Append(Route{Path: "/posts/:id"})
  router.Append(Route{Path: "/comments/:id"})

  route, vars := router.ParseRoute("/posts/1")

  if route == nil {
    fmt.Println("ERROR")
  }

  if vars[":id"] != "1" {
    fmt.Println("ERROR")
  }

  route, _ = router.ParseRoute("/not/in/router")

  if route != nil {
    fmt.Println("ERROR")
  }
}
