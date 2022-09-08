package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kapralovs/warehouse/internal/auth"
	"github.com/kapralovs/warehouse/internal/models"
	"github.com/kapralovs/warehouse/internal/users"
)

type Handler struct {
	usecase auth.UseCase
}

func NewHandler(uc auth.UseCase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) SignUp() func(http.ResponseWriter, *http.Request) {
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

		newUser := &models.User{}
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

		log.Printf("New profile created by user \"%s\"\n", user.Username)
		s.storage.SaveUser(newUser)

		fmt.Fprintln(w, "User profile is created!")
	}
}

func (h *Handler) SignIn() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if username, password, ok := r.BasicAuth(); ok {
			err := h.usecase.SignIn(username, password)
			if err != nil {

			}
			// if profile, ok := ds.Profiles[username]; ok {

			// if profile.Account.Password == password {
			// 	ds.Profiles[username].Account.IsOnline = true
			// 	fmt.Println("Profile pointer", profile)                             //DEBUG
			// 	fmt.Println("ds.Profiles[username] pointer", ds.Profiles[username]) //DEBUG
			// 	log.Printf("User \"%s\" is authorised\n", profile.Account.Username)
			// 	return profile, nil
			// }
			// }
			// if user, ok := checkCredentials(ds, username, password); ok {
			// 	log.Printf("User \"%s\" is authorised\n", user.Account.Username)
			// 	return user, nil
			// }
		}

		// w.Header().Add("WWW-Authenticate", "Basic realm="+r.Header.Get("Authorization"))
		// w.WriteHeader(http.StatusUnauthorized)

		// return nil, auth.ErrAuthFailed
	}
}

// func authorization(ds *data.DataStorage, w http.ResponseWriter, r *http.Request) (*models.User, error) {

// 	if username, password, ok := r.BasicAuth(); ok {
// 		if profile, ok := ds.Profiles[username]; ok {
// 			if profile.Account.Password == password {
// 				ds.Profiles[username].Account.IsOnline = true
// 				fmt.Println("Profile pointer", profile)                             //DEBUG
// 				fmt.Println("ds.Profiles[username] pointer", ds.Profiles[username]) //DEBUG
// 				log.Printf("User \"%s\" is authorised\n", profile.Account.Username)
// 				return profile, nil
// 			}
// 		}
// 		// if user, ok := checkCredentials(ds, username, password); ok {
// 		// 	log.Printf("User \"%s\" is authorised\n", user.Account.Username)
// 		// 	return user, nil
// 		// }
// 	}

// 	w.Header().Add("WWW-Authenticate", "Basic realm="+r.Header.Get("Authorization"))
// 	w.WriteHeader(http.StatusUnauthorized)

// 	return nil, auth.ErrAuthFailed
// }
