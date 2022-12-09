package http

import (
	"fmt"
	q "miracle-calendar/db"
	"net/http"
)


func MainGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Miracle Calendar"))
	q.GetMain()
}

func MainPostHandler(w http.ResponseWriter, r *http.Request) {
	var routine, memo string
	fmt.Println("Write Routine")
	fmt.Scan(&routine)
	fmt.Println("Write Memo")
	fmt.Scan(&memo)
	
	q.PostMain(routine, memo)
}

func MainPutHandler(w http.ResponseWriter, r *http.Request) {
	var routine, memo string
	var is_achieved, id int
	fmt.Println("Write id for updating")
	fmt.Scan(&id)
	fmt.Println("Update routine & memo & is_achieve")
	fmt.Scan(&routine, &memo, &is_achieved)

	q.PutMain(routine, memo, is_achieved, id)
}

func MainDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Miracle Calendar"))

	q.DeleteMain()
}