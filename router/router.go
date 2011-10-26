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

func (r *Router) AddController(controller Controller) {
  routes := controller.Routes()

  for i := 0; i < len(routes); i++ {
    WebRouter.Append(routes[i])
  }
}
