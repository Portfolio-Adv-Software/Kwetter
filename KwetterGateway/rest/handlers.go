package rest

import "net/http"

func HomePageFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//handle stuff
	}
}

func FetchAllKwettersFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//handle stuff
	}
}
func FetchKwetterFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//handle stuff
	}
}

func CreateKwetterFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//handle stuff
	}
}
