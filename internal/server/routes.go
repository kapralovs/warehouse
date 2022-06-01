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
		if user, err := authorization(s.storage, w, r); err != nil {

		}
	}
}

//--------------------Tasks--------------------

func (s *server) createTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) editTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) deleteTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) getTasks() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) geTaskByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

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

//--------------------Products--------------------

func (s *server) createProduct() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) editProduct() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) deleteProduct() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) getProducts() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) getProductByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

//--------------------Cells--------------------

func (s *server) createCell() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) editCell() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) deleteCell() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) getCells() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) getCellByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
