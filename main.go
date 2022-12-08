package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Miracle Calendar"))
}

func dailyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Check your daily routine"))
}

func dailyMemoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Write memo"))
}

func main() {
	mux := chi.NewRouter()
	mux.Get("/", calendarHandler)
	mux.Get("/daily", dailyHandler)
	mux.Get("/daily/memo", dailyMemoHandler)
	http.ListenAndServe(":8080", mux)
}