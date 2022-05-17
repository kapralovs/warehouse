package data

import (
	"fmt"

	"github.com/kapralovs/warehouse/internal/users"
)

func CreateUser() error {
	return nil
}

func (ds *DataStorage) LoadUser(id string) (*users.Profile, error) {
	if profile, ok := ds.Profiles[id]; ok {
		return profile, nil
	}

	return nil, fmt.Errorf("user with this id does not exist", id)
}

func DeleteUser() error {
	return nil
}

func (ds *DataStorage) SaveUser(up *users.Profile) error {

	return nil
}
