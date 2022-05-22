package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapralovs/warehouse/internal/data"
)

func New() *server {
	return &server{
		router: mux.NewRouter(),
	}
}

func (s *server) initRouter() {
	// s.router.HandleFunc("/", s.homepage())
	//Tasks
	s.router.HandleFunc("/task", s.getUsers()).Methods("GET")
	s.router.HandleFunc("/task", s.createUser()).Methods("POST")
	s.router.HandleFunc("/task/{id}", s.editUser()).Methods("POST")
	s.router.HandleFunc("/task/{id}", s.deleteUser()).Methods("DELETE")
	s.router.HandleFunc("/task/{id}", s.getUserByID()).Methods("GET")
	// s.router.HandleFunc("/task/start/{id}", s.startTask())
	// s.router.HandleFunc("/task/complete/{id}", s.completeTask())
	//Users
	s.router.HandleFunc("/user", s.getUsers()).Methods("GET")
	s.router.HandleFunc("/user", s.createUser()).Methods("POST")
	s.router.HandleFunc("/user/{id}", s.editUser()).Methods("POST")
	s.router.HandleFunc("/user/{id}", s.deleteUser()).Methods("DELETE")
	s.router.HandleFunc("/user/{id}", s.getUserByID()).Methods("GET")
	//Products
	s.router.HandleFunc("/product", s.getUsers()).Methods("GET")
	s.router.HandleFunc("/product", s.createUser()).Methods("POST")
	s.router.HandleFunc("/product/{id}", s.editUser()).Methods("POST")
	s.router.HandleFunc("/product/{id}", s.deleteUser()).Methods("DELETE")
	s.router.HandleFunc("/product/{id}", s.getUserByID()).Methods("GET")
	//Cells
	s.router.HandleFunc("/cell", s.getUsers()).Methods("GET")
	s.router.HandleFunc("/cell", s.createUser()).Methods("POST")
	s.router.HandleFunc("/cell/{id}", s.editUser()).Methods("POST")
	s.router.HandleFunc("/cell/{id}", s.deleteUser()).Methods("DELETE")
	s.router.HandleFunc("/cell/{id}", s.getUserByID()).Methods("GET")
}

func (s *server) Run() error {
	s.initRouter()
	st := data.NewStorage()
	st.Init()
	s.storage = st
	log.Println("Server starting on port 8080...")
	return http.ListenAndServe(":8080", s.router)
}
