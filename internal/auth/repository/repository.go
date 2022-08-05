package repository

import (
	"sync"

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
