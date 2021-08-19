package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    router.
        Name(route.Name).
        Methods(route.Method).
        Path(route.Pattern).
        Handler(route.HandlerFunc)
  }
  return router
}

var routes = Routes{
  Route{
    "HealthCheck",
    "GET",
    "/healthCheck",
    HealthCheck,
  },
  Route{
      "PostSatellites",
      "POST",
      "/topsecret",
      PostSatellites,
  },
  Route{
      "PostSatellitesSplit",
      "POST",
      "/topsecret_split/{name_satellite}",
      PostSatellitesSplit,
    },
  Route{
      "GetSatellitesSplit",
      "GET",
      "/topsecret_split/",
      GetSatellitesSplit,
    },
}
