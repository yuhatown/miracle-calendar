package main

import (
	"net/http"
	"miracle-calendar/model"
	"github.com/go-chi/chi"
)


func main() {
	mux := chi.NewRouter()
	mux.Get("/routine", model.getRoutineHandler)
	mux.Post("/routine", model.postRoutineHandler)
	mux.Put("/routine/:id", model.putRoutineHandler)
	mux.Delete("routine/:id", model.deleteRoutineHandler)

	http.ListenAndServe(":8080", mux)	
}