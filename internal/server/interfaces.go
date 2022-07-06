package server

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/data"
)

// Server structure
type server struct {
	router  *mux.Router
	storage *data.DataStorage
}
