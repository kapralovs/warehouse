package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/kapralovs/warehouse/internal/data"
	"github.com/kapralovs/warehouse/internal/users"
)

func checkCredentials(ds *data.DataStorage, username, password string) (*users.Profile, bool) {
	for _, profile := range ds.Profiles {
		if profile.Account.Username == username && profile.Account.Password == password {
			return profile, true
		}
	}

	return nil, false
	// errors.New("authorisation failed because credentials are incorrect")
}

func authorization(ds *data.DataStorage, w http.ResponseWriter, r *http.Request) (*users.Profile, error) {

	if username, password, ok := r.BasicAuth(); ok {
		if user, ok := checkCredentials(ds, username, password); ok {
			log.Printf("User \"%s\" is authorised\n", user.Account.Username)
			return user, nil
		}
	}

	w.Header().Add("WWW-Authenticate", "Basic realm="+r.Header.Get("Authorization"))
	w.WriteHeader(http.StatusUnauthorized)

	return nil, errors.New("authorization failed")
}
