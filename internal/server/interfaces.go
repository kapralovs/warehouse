package server

import (
	"github.com/gorilla/mux"
)

type server struct {
	router  *mux.Router
	storage *datastorage.DataStorage
}
