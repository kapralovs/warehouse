package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoints(r *mux.Router, uc warehouse.UseCase) {
	warehouseHandler := NewHandler(uc)

	r.HandleFunc("/cell", warehouseHandler.getCell()).Methods(http.MethodGet)
	r.HandleFunc("/cell", warehouseHandler.createCell()).Methods(htyp.MethodPost)
	r.HandleFunc("/cell/{id}", warehouseHandler.editCell()).Methods(http.MethodPost)
	r.HandleFunc("/cell/{id}", warehouseHandler.deleteCell()).Methods(http.MethodDelete)
	r.HandleFunc("/cell/{id}", warehouseHandler.getCellByID()).Methods(http.MethodGet)
}
