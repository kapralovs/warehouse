package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func New() *server {
	return &server{
		router: mux.NewRouter(),
	}
}

func (s *server) Run() error {
	// s.initRouter()
	// st := data.NewStorage()
	// st.Init()
	// s.storage = st
	log.Println("Server starting on port 8080...")
	return http.ListenAndServe(":8080", s.router)
}
