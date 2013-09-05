package home_controller

import (
  "net/http"
  "io"
  "../router"
)

type Controller struct { }

func (hc *Controller) Routes() []router.Route {
  routes := []router.Route{ router.Route{Path: "/posts/:id", Handler: HomePage} }
  return routes
}

func HomePage(writer http.ResponseWriter, r *http.Request, vars map[string]string) {
  io.WriteString(writer, "Found id " + vars[":id"])
}
