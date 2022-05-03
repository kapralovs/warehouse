package server

import "github.com/gorilla/mux"

func New() *server {
	return &server{
		router: mux.NewRouter(),
	}
}
