package main

import (
	"miracle-calendar/model"
	"miracle/miracle-calendar/model"
	"net/http"

	"github.com/go-chi/chi"
)


func main() {
	mux := chi.NewRouter()
	mux.Get("/routine", model.getRoutineHandler)
	mux.Post("/routine", model.postRoutineHandler)
	mux.Put("/routine/:id", model.putRoutineHandler)
	mux.Delete("routine/:id", model.deleteRoutineHandler)

	mux.Get("/user", model.GetUser)
	mux.Post("/user", model.PostUser)
	mux.Put("/user/:id", model.PutUser)
	mux.Delete("/user/:id", model.DeleteUser)

	http.ListenAndServe(":8080", mux)	
}