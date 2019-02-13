package router

import (
	"Books_A_Hundreds/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var c = &controller.Controller{}

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: c.HomePageFunc,
	},
	Route{
		Name:        "GetAllBooks",
		Method:      "GET",
		Pattern:     "/all",
		HandlerFunc: c.ReturnAllBooks,
	},
	Route{
		Name:        "GetSingleBook",
		Method:      "GET",
		Pattern:     "/book/{id}",
		HandlerFunc: c.ReturnSingleBook,
	},
}

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		fmt.Println("Adding route: ", route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

	return router
}
