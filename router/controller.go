package router

type Controller interface {
  Routes() []Route
}
