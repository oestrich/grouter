package router

import (
  "io"
  "http"
)

func RoutePath(writer http.ResponseWriter, r *http.Request) {
  io.WriteString(writer, "Hit a page\n")
  io.WriteString(writer, r.Method + "\n")
  io.WriteString(writer, r.URL.Raw + "\n")
}

func SetupRouter() {
  http.HandleFunc("/", RoutePath)
}
