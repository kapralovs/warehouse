package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kapralovs/warehouse/internal/users"
)

type Handler struct {
	usecase users.UseCase
}

func NewHandler(uc users.UseCase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) editUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) deleteUser() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) getUsers() func(http.ResponseWriter, *http.Request) {
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

func (h *Handler) getUserByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
