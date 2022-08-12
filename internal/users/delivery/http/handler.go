package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *server) editUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) deleteUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) getUsers() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := authorization(s.storage, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		for id := range s.storage.Profiles {
			profile, err := s.storage.LoadUser(id)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			jsonAsBytes, err := json.Marshal(profile)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintln(w, string(jsonAsBytes))
		}
	}
}

func (s *server) getUserByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
