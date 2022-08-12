package repository

import (
	"errors"
	"log"
	"sync"

	"github.com/kapralovs/warehouse/internal/models"
	"github.com/kapralovs/warehouse/internal/users"
)

type UserRepository struct {
	mu    *sync.Mutex
	users map[string]*models.User
}

func (ur *UserRepository) LoadUser(id string) (*models.User, error) {
	if profile, ok := ur.users[id]; ok {
		return profile, nil
	}

	return nil, errors.New("user with this id does not exist")
}

func (ur *UserRepository) DeleteUser(id string) error {
	if _, ok := ur.users[id]; ok {
		delete(ur.users, id)
		return nil
	}

	return errors.New("it is not possible to delete a user profile because it does not exist")
}

func (ur *UserRepository) SaveUser(u *models.User) error {
	if u != nil {
		if u.ID == "" {
			return errors.New("can't save profile with empty ID field")
		}

		ur.mu.Lock()
		defer ur.mu.Unlock()
		ur.users[u.ID] = u

		log.Printf("The profile \"%s\" is saved.\n", ur.users[u.ID].Username)
		return nil
	}

	return errors.New("can't save current profile, because profile is nil")
}

func (ur *UserRepository) CheckProfileDuplicates(u *models.User) error {
	if u != nil {
		if u.ID != "" {
			for id, user := range ur.users {
				if id == u.ID {
					return errors.New("profile with this ID already exists")
				}
				if user.Username == u.Username {
					return errors.New("profile with this username already exists")
				}

				return nil
			}

			log.Println()
			return errors.New("empty profile ID")
		}
	}

	return errors.New("edited user profile is nil")
}

// Инициализация хранилища (DEBUG!!! Позже удалить)
func (ur *UserRepository) Init() {
	ur.users["admin"] = &models.User{
		Lastname: "Adminov",
		Name:     "Admin",
		Age:      36,
		Role:     users.Manager,
		Salary:   50000,
		ID:       "1",
		Username: "admin",
		Password: "admin",
		IsAdmin:  true,
	}
}
