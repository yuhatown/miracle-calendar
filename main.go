package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Miracle Calendar"))
}

func routineHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Check your daily routine"))
}

func routineMemoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Write memo"))
}

func main() {

	mux := chi.NewRouter()
	mux.Get("/", calendarHandler)
	mux.Get("/daily", routineHandler)
	mux.Get("/daily/memo", routineMemoHandler)
	http.ListenAndServe(":8080", mux)


	err := godotenv.Load(".env")
	if err != nil {
  		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}