package http

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/products"
)

func RegisterEndpoints(r *mux.Router, uc products.UseCase) {
	h := NewHandler(uc)
	// r.HandleFunc("/user", h.getUsers()).Methods("GET")
	// r.HandleFunc("/user", h.createUser()).Methods("POST")
	r.HandleFunc("/user/{id}", h.editUser()).Methods("POST")
	r.HandleFunc("/user/{id}", h.deleteUser()).Methods("DELETE")
	// r.HandleFunc("/user/{id}", h.getUserByID()).Methods("GET")
}
