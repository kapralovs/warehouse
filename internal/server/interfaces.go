package server

import (
	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/storage"
)

type server struct {
	router  *mux.Router
	storage *storage.Storage
}
