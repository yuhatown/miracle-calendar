package app

import (
	"fmt"
	"miracle-calendar/model"
	"net/http"
)

type routine struct {

}

type user struct {

}

func (a routine) getRoutineHandler(w http.ResponseWriter, r *http.Request) {
	model.GetRoutine()
}

func (a routine) postRoutineHandler(w http.ResponseWriter, r *http.Request) {
	var routine, memo string
	fmt.Println("Write Routine")
	fmt.Scan(&routine)
	fmt.Println("Write Memo")
	fmt.Scan(&memo)
	
	model.PostRoutine(routine, memo)
}

func (a routine) putRoutineHandler(w http.ResponseWriter, r *http.Request) {
	var routine, memo string
	var is_achieved, id int
	fmt.Println("Write id for updating")
	fmt.Scan(&id)
	fmt.Println("Update routine & memo & is_achieve")
	fmt.Scan(&routine, &memo, &is_achieved)

	model.PutRoutine(routine, memo, is_achieved, id)
}

func (a routine) deleteRoutineHandler(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Println("Write id for deleting")
	fmt.Scan(&id)

	model.DeleteRoutine(id)
}

func (u user) getUserHandler(w http.ResponseWriter, r *http.Request) {
	model.GetUser()
}

func (u user) postUserHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	fmt.Println("Write Name")
	fmt.Scan(&name)
	
	model.PostUser(name)
}

func (u user) putUserHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	var id int
	fmt.Println("Write id for updating")
	fmt.Scan(&id)
	fmt.Println("Update routine & memo & is_achieve")
	fmt.Scan(&name)

	model.PutUser(name, id)
}

func (u user) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Println("Write id for deleting")
	fmt.Scan(&id)

	model.DeleteUser(id)
}