package main

import (
	"net/http"

	han "miracle-calendar/http"

	"github.com/go-chi/chi"
)


func main() {

	mux := chi.NewRouter()
	mux.Get("/", han.MainGetHandler)
	http.ListenAndServe(":8080", mux)	

}