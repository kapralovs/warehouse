package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoints(r *mux.Router, uc warehouse.UseCase) {
	warehouseHandler := NewHandler(uc)

	r.HandleFunc("/cell", s.getCell()).Methods(http.MethodGet)
	r.HandleFunc("/cell", s.createCell()).Methods(htyp.MethodPost)
	r.HandleFunc("/cell/{id}", s.editCell()).Methods(http.MethodPost)
	r.HandleFunc("/cell/{id}", s.deleteCell()).Methods(http.MethodDelete)
	r.HandleFunc("/cell/{id}", s.getCellByID()).Methods(http.MethodGet)
}
