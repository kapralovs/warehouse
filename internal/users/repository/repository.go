package repository

import (
	"errors"
	"log"
	"sync"

	"github.com/kapralovs/warehouse/internal/models"
)

type UserRepository struct {
	mu       *sync.Mutex
	Profiles map[string]*models.Profile
}

func (ur *UserRepository) LoadUser(id string) (*models.Profile, error) {
	if profile, ok := ur.Profiles[id]; ok {
		return profile, nil
	}

	return nil, errors.New("user with this id does not exist")
}

func (ur *UserRepository) DeleteUser(id string) error {
	if _, ok := ur.Profiles[id]; ok {
		delete(ur.Profiles, id)
		return nil
	}

	return errors.New("it is not possible to delete a user profile because it does not exist")
}

func (ur *UserRepository) SaveUser(p *models.Profile) error {
	if p != nil {
		if p.Account.ID == "" {
			return errors.New("can't save profile with empty ID field")
		}

		ur.mu.Lock()
		defer ur.mu.Unlock()
		ur.Profiles[p.Account.ID] = p

		log.Printf("The profile \"%s\" is saved.\n", ur.Profiles[p.Account.ID].Account.Username)
		return nil
	}

	return errors.New("can't save current profile, because profile is nil")
}
