package main

import (
	"miracle-calendar/app"
	"net/http"

	"github.com/go-chi/chi"
)


func main() {
	mux := chi.NewRouter()
	mux.Get("/routine", app.getRoutineHandler)
	mux.Post("/routine", app.postRoutineHandler)
	mux.Put("/routine/:id", app.putRoutineHandler)
	mux.Delete("routine/:id", app.deleteRoutineHandler)

	mux.Get("/user", app.getUserHandler)
	mux.Post("/user", app.postUserHandler)
	mux.Put("/user/:id", app.putUserHandler)
	mux.Delete("/user/:id", app.deleteUserHandler)

	http.ListenAndServe(":8080", mux)	
}