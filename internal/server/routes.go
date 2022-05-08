package server

import "net/http"

func (s *server) homepage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

//--------------------Tasks--------------------

func (s *server) completeTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

//--------------------Users--------------------

func (s *server) createUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

//--------------------Products--------------------

//--------------------Cells--------------------
