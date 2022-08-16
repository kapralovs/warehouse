package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/warehouse"
)

func RegisterHTTPEndpoints(r *mux.Router, uc warehouse.UseCase) {
	wh := NewHandler(uc)

	r.HandleFunc("/cell", wh.getCells()).Methods(http.MethodGet)
	r.HandleFunc("/cell", wh.createCell()).Methods(http.MethodPost)
	r.HandleFunc("/cell/{id}", wh.editCell()).Methods(http.MethodPost)
	r.HandleFunc("/cell/{id}", wh.deleteCell()).Methods(http.MethodDelete)
	r.HandleFunc("/cell/{id}", wh.getCellByID()).Methods(http.MethodGet)
}
