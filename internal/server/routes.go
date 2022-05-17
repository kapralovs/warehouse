package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kapralovs/warehouse/internal/users"
)

func (s *server) homepage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

//--------------------Tasks--------------------

func (s *server) startTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) completeTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

//--------------------Users--------------------

func (s *server) createUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := authorization(s.storage, w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if err := users.CheckAdminRights(user); err != nil {
			fmt.Fprintln(w, err)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, err)
			return
		}

		newUser := &users.Profile{}
		if err := json.Unmarshal(body, &newUser); err != nil {
			log.Println(err)
			fmt.Fprintln(w, err)
			return
		}

		if err := s.storage.CheckProfileDuplicates(newUser); err != nil {
			log.Println(err)
			fmt.Fprintln(w, err)
			return
		}

		log.Printf("New profile created by user \"%s\"\n", user.Account.Username)
		s.storage.SaveUser(newUser)

		fmt.Fprintln(w, "User profile is created!")
	}
}

func (s *server) deleteUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

//--------------------Products--------------------

//--------------------Cells--------------------
