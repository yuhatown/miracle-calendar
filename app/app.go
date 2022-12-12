package app

import (
	"fmt"
	q "miracle-calendar/model"
	"net/http"
)

type routine struct {

}


func (a *routine) getRoutineHandler(w http.ResponseWriter, r *http.Request) {
	q.GetRoutine()
}

func (a *routine) postRoutineHandler(w http.ResponseWriter, r *http.Request) {
	var routine, memo string
	fmt.Println("Write Routine")
	fmt.Scan(&routine)
	fmt.Println("Write Memo")
	fmt.Scan(&memo)
	
	q.PostRoutine(routine, memo)
}

func putRoutineHandler(w http.ResponseWriter, r *http.Request) {
	var routine, memo string
	var is_achieved, id int
	fmt.Println("Write id for updating")
	fmt.Scan(&id)
	fmt.Println("Update routine & memo & is_achieve")
	fmt.Scan(&routine, &memo, &is_achieved)

	q.PutRoutine(routine, memo, is_achieved, id)
}

func deleteRoutineHandler(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Println("Write id for deleting")
	fmt.Scan(&id)

	q.DeleteRoutine(id)
}
