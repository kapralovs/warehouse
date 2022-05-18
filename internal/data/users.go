package data

import (
	"errors"
	"log"

	"github.com/kapralovs/warehouse/internal/users"
)

func (ds *DataStorage) LoadUser(id string) (*users.Profile, error) {
	if profile, ok := ds.Profiles[id]; ok {
		return profile, nil
	}

	return nil, errors.New("user with this id does not exist")
}

func (ds *DataStorage) DeleteUser(id string) error {
	if _, ok := ds.Profiles[id]; ok {
		delete(ds.Profiles, id)
		return nil
	}

	return errors.New("it is not possible to delete a user profile because it does not exist")
}

func (ds *DataStorage) SaveUser(p *users.Profile) error {
	if p != nil {
		if p.Account.ID == "" {
			return errors.New("can't save profile with empty ID field")
		}

		ds.mu.Lock()
		defer ds.mu.Unlock()
		ds.Profiles[p.Account.ID] = p

		log.Printf("The profile \"%s\" is saved.\n", ds.Profiles[p.Account.ID].Account.Username)
		return nil
	}

	return errors.New("can't save current profile, because profile is nil")
}
