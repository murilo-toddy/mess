package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const (
	baseTodoRoute = "todos"
)

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", `/` + baseTodoRoute, TodoIndex},
	Route{"TodoShow", "GET", `/` + baseTodoRoute + `/{todoId}`, TodoShow},
	Route{"TodoCreate", "POST", `/` + baseTodoRoute, TodoCreate},
}
