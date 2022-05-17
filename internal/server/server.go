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

func (s *server) initRouter() {
	s.router.HandleFunc("/", s.homepage())
	s.router.HandleFunc("/task/start/{id}", s.startTask())
	s.router.HandleFunc("/task/complete/{id}", s.completeTask())
	s.router.HandleFunc("/user/create", s.createUser())
	s.router.HandleFunc("/user/edit/{id}", s.deleteUser())
}

func (s *server) Run() error {
	s.initRouter()
	log.Println("Server starting on port 8080...")
	return http.ListenAndServe(":8080", s.router)
}
