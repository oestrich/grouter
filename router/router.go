package router

import (
  "io"
  "http"
  "strings"
)

type Route struct {
  Path string
}

func RoutePath(writer http.ResponseWriter, r *http.Request) {
  io.WriteString(writer, "Hit a page\n")
  io.WriteString(writer, r.Method + "\n")
  io.WriteString(writer, r.URL.Raw + "\n")
}

func SetupRouter() {
  http.HandleFunc("/", RoutePath)
}

func (r *Route) Split() []string {
  return strings.Split(r.Path, "/")
}

func (r *Route) ParseRoute(path string) (bool, string, map[string]string) {
  vars := make(map[string]string)
  sections_route := r.Split()
  sections_path := strings.Split(path, "/")

  if len(sections_route) != len(sections_path) {
    return false, "", vars
  }

  for i := 0; i < len(sections_route); i++ {
    if sections_route[i] == "" {
      continue
    }

    if strings.HasPrefix(sections_route[i], ":") {
      if i < len(sections_path) {
        vars[sections_route[i]] = sections_path[i]
      }
    }
  }

  return true, "handler", vars
}
