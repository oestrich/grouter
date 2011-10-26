package router

import (
  "io"
  "http"
  "strings"
)

type Route struct {
  Path string
}

type Router struct {
  Routes []Route
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

func (router *Router) Append(route Route) {
  router.Routes = append(router.Routes, route)
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

    switch {
    case strings.HasPrefix(sections_route[i], ":"):
      vars[sections_route[i]] = sections_path[i]
    case sections_route[i] != sections_path[i]:
      return false, "", vars
    }
  }

  return true, "handler", vars
}

func (r *Router) ParseRoute(path string) (*Route, map[string]string) {
  for i := 0; i < len(r.Routes); i++ {
    pass, _, vars := r.Routes[i].ParseRoute(path)

    if pass {
      return &r.Routes[i], vars
    }
  }

  return nil, nil
}
