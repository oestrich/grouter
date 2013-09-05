package router

import (
  "io"
  "net/http"
  "fmt"
)

var WebRouter = &Router{}

func RoutePath(writer http.ResponseWriter, r *http.Request) {
  fmt.Println("Accessing page " + r.URL.Path)

  route, vars := WebRouter.ParseRoute(r.URL.Path)

  if route != nil {
    route.Handler(writer, r, vars)
  } else {
    writer.WriteHeader(http.StatusNotFound)
    io.WriteString(writer, "404 Page Not Found")
  }
}

func SetupRouter() {
  http.HandleFunc("/", RoutePath)
}
