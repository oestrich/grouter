package router

import (
  "io"
  "http"
  "fmt"
)

var WebRouter = &Router{}

func RoutePath(writer http.ResponseWriter, r *http.Request) {
  fmt.Println("Accessing page " + r.URL.Raw)

  route, vars := WebRouter.ParseRoute(r.URL.Raw)

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
