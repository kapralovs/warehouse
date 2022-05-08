package server

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/data"
)

type server struct {
	router  *mux.Router
	storage *data.DataStorage
}
