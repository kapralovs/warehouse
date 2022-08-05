package server

import (
	"github.com/gorilla/mux"
)

// Server structure
type server struct {
	router *mux.Router
	// storage *data.DataStorage
}
