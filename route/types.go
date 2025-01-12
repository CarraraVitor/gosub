package route

import "net/http"

type Router struct {
    Routes []Route
}

type Route struct {
    Path string
    Handler func(w http.ResponseWriter, r *http.Request)
}

func (r *Router) Register(route Route) {
    r.Routes = append(r.Routes, route)
}

