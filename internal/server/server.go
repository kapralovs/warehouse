package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() *server {
	return &server{
		router: mux.NewRouter(),
	}
}

func (s *server) initRouter() {
	s.router.HandlerFunc("/")
}

func (s *server) Run() error {
	return http.ListenAndServe(":8080", s.router)
}
