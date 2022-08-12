package repository

import (
	"sync"

	"github.com/kapralovs/warehouse/internal/auth"
	"github.com/kapralovs/warehouse/internal/models"
)

type AuthRepository struct {
	mu    *sync.Mutex
	users map[string]*models.User
}

func New() *AuthRepository {
	return &AuthRepository{
		mu:    new(sync.Mutex),
		users: make(map[string]*models.User),
	}
}

func (ar *AuthRepository) CreateUser(user *models.User) error {
	ar.mu.Lock()
	ar.users[user.ID] = user
	ar.mu.Unlock()
	return nil
}

func (ar *AuthRepository) GetUser(username, password string) (*models.User, error) {
	for _, u := range ar.users {
		if u.Username == username && u.Password == password {
			return u, nil
		}
	}

	return nil, auth.ErrUserNotFound
}

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
