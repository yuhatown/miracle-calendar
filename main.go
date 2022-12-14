package main

import (
	"miracle-calendar/model"
	"miracle/miracle-calendar/model"
	"net/http"

	"github.com/go-chi/chi"
)


func main() {
	mux := chi.NewRouter()
	mux.Get("/routine", model.GetRoutine)
	mux.Post("/routine", model.PostRoutine)
	mux.Put("/routine/:id", model.PutRoutine)
	mux.Delete("routine/:id", model.DeleteRoutine)

	mux.Get("/user", model.GetUser)
	mux.Post("/user", model.PostUser)
	mux.Put("/user/:id", model.PutUser)
	mux.Delete("/user/:id", model.DeleteUser)

	http.ListenAndServe(":8080", mux)	
}