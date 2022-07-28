package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/kapralovs/warehouse/internal/data"
	"github.com/kapralovs/warehouse/internal/models"
)

// func checkCredentials(ds *data.DataStorage, username, password string) (*users.Profile, bool) {
// 	for id := range ds.Profiles {
// 		if ds.Profiles[id].Account.Username == username && ds.Profiles[id].Account.Password == password {
// 			ds.Profiles[id].Account.IsOnline = true
// 			return ds.Profiles[id], true
// 		}
// 	}

// 	return nil, false
// 	// errors.New("authorisation failed because credentials are incorrect")
// }

// Авторизация пользователя
func authorization(ds *data.DataStorage, w http.ResponseWriter, r *http.Request) (*models.Profile, error) {

	if username, password, ok := r.BasicAuth(); ok {
		if profile, ok := ds.Profiles[username]; ok {
			if profile.Account.Password == password {
				ds.Profiles[username].Account.IsOnline = true
				fmt.Println("Profile pointer", profile)                             //DEBUG
				fmt.Println("ds.Profiles[username] pointer", ds.Profiles[username]) //DEBUG
				log.Printf("User \"%s\" is authorised\n", profile.Account.Username)
				return profile, nil
			}
		}
		// if user, ok := checkCredentials(ds, username, password); ok {
		// 	log.Printf("User \"%s\" is authorised\n", user.Account.Username)
		// 	return user, nil
		// }
	}

	w.Header().Add("WWW-Authenticate", "Basic realm="+r.Header.Get("Authorization"))
	w.WriteHeader(http.StatusUnauthorized)

	return nil, errors.New("authorization failed")
}
