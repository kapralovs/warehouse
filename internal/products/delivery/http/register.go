package http

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/products"
)

func RegisterEndpoints(r *mux.Router, uc products.UseCase) {
	h := NewHandler(uc)
	r.HandleFunc("/product", h.getProducts()).Methods("GET")
	r.HandleFunc("/product", h.createProduct()).Methods("POST")
	r.HandleFunc("/product/{id}", h.editProduct()).Methods("POST")
	r.HandleFunc("/product/{id}", h.deleteProduct()).Methods("DELETE")
	r.HandleFunc("/product/{id}", h.getProductByID()).Methods("GET")
}
