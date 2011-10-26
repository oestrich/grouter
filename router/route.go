package router

import (
  "strings"
  "http"
)

type Route struct {
  Path string
  Handler func (http.ResponseWriter, *http.Request, map[string]string)
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

    switch {
    case strings.HasPrefix(sections_route[i], ":"):
      vars[sections_route[i]] = sections_path[i]
    case sections_route[i] != sections_path[i]:
      return false, "", vars
    }
  }

  return true, "handler", vars
}
