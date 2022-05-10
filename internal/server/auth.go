package server

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/kapralovs/warehouse/internal/data"
	"github.com/kapralovs/warehouse/internal/users"
)

func checkCredentials(ds *data.DataStorage, encodedCreds string) (*users.Profile, error) {
	decodedCreds, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		return nil, err
	}

	creds := strings.Split(string(decodedCreds), ":")

	for _, profile := range ds.Profiles {
		if profile.Account.Username == creds[0] && profile.Account.Password == creds[1] {
			log.Printf("Credentials \"%s\" are checked\n", encodedCreds)
			return profile, nil
		}
	}

	return nil, errors.New("authorisation failed because credentials are incorrect")
}

func authorization(ds *data.DataStorage, w http.ResponseWriter, r *http.Request) (*users.Profile, error) {
	headerValue := r.Header.Get("Authorization")
	encodedCreds := headerValue[len("Basic "):]
	user, err := checkCredentials(ds, encodedCreds)
	if err != nil {
		w.Header().Add("WWW-Authenticate", "Basic realm="+encodedCreds)
		w.WriteHeader(http.StatusUnauthorized)

		return nil, errors.New("authorization failed")
	}

	log.Printf("User \"%s\" is authorised\n", user.Account.Username)
	return user, nil
}
