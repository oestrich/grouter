package router

import (
  "io"
  "http"
)

var WebRouter = &Router{}

func RoutePath(writer http.ResponseWriter, r *http.Request) {
  route, vars := WebRouter.ParseRoute(r.URL.Raw)

  if route != nil {
    io.WriteString(writer, "Matched a Route!\n")

    for k, v := range vars {
      io.WriteString(writer, k + " => " + v + "\n")
    }

    route.Handler(writer, r, vars)
  } else {
    writer.WriteHeader(http.StatusNotFound)
    io.WriteString(writer, "404 Page Not Found")
  }
}

func SetupRouter() {
  http.HandleFunc("/", RoutePath)
}
