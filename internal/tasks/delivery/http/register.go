package http

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/tasks"
)

func RegisterEndpoints(r *mux.Router, uc tasks.UseCase) {
	t := New(uc)

	r.HandleFunc("/task", t.createTask()).Methods("GET")
	r.HandleFunc("/task", t.editTask()).Methods("POST")
	r.HandleFunc("/task/{id}", t.deleteTask()).Methods("POST")
	r.HandleFunc("/task/{id}", t.getTasks()).Methods("DELETE")
	r.HandleFunc("/task/{id}", t.geTaskByID()).Methods("GET")
	r.HandleFunc("/task/start/{id}", t.startTask())
	r.HandleFunc("/task/complete/{id}", t.completeTask())
}
