package kwetter

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Name).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func InitRouter() {
	routes := Routes{
		Route{
			Name:        "Home",
			Method:      "GET",
			HandlerFunc: HomePageFunc(),
		},
		Route{
			Name:        "/kwetter/{id}",
			Method:      "GET",
			HandlerFunc: FetchKwetterFunc(),
		},
		Route{
			Name:        "/kwetter/all",
			Method:      "GET",
			HandlerFunc: FetchAllKwettersFunc(),
		},
		Route{
			Name:        "/create",
			Method:      "POST",
			HandlerFunc: CreateKwetterFunc(),
		},
	}

	router := NewRouter(routes)
	http.ListenAndServe(":8080", router)
}
