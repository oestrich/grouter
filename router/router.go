package router

type Router struct {
  Routes []Route
}

func (router *Router) Append(route Route) {
  router.Routes = append(router.Routes, route)
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
