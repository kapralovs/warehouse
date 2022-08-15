package http

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/products"
)

func RegisterEndpoints(r *mux.Router, uc products.UseCase) {
	ph := NewHandler(uc7)
	r.HandleFunc("/product", ph.getUsers()).Methods("GET")
	r.HandleFunc("/product", ph.createUser()).Methods("POST")
	r.HandleFunc("/product/{id}", ph.editUser()).Methods("POST")
	r.HandleFunc("/product/{id}", ph.deleteUser()).Methods("DELETE")
	r.HandleFunc("/product/{id}", ph.getUserByID()).Methods("GET")
}
