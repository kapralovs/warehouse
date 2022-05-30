package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/kapralovs/warehouse/internal/data"
	"github.com/kapralovs/warehouse/internal/users"
)

func checkCredentials(ds *data.DataStorage, username, password string) bool {
	/*
		decodedCreds, err := base64.StdEncoding.DecodeString(encodedCreds)
		if err != nil {
			return nil, err
		}

		creds := strings.Split(string(decodedCreds), ":")
	*/

	for _, profile := range ds.Profiles {
		if profile.Account.Username == username && profile.Account.Password == password {
			return true
		}
	}

	return false
	// errors.New("authorisation failed because credentials are incorrect")
}

func authorization(ds *data.DataStorage, w http.ResponseWriter, r *http.Request) (*users.Profile, error) {
	/*
		headerValue := r.Header.Get("Authorization")
		encodedCreds := headerValue[len("Basic "):]
	*/
	username, password, ok := r.BasicAuth()
	if ok && checkCredentials(ds, username, password) {
		return
	}
	if err != nil {
		w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds)
		w.WriteHeader(http.StatusUnauthorized)

		return nil, errors.New("authorization failed")
	}

	log.Printf("User \"%s\" is authorised\n", user.Account.Username)
	return user, nil
}
