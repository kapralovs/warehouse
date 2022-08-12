package usecase

import (
	"github.com/kapralovs/warehouse/internal/auth"
	"github.com/kapralovs/warehouse/internal/models"
)

type AuthUseCase struct {
	repo auth.UserRepository
}

func New(userRepo auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		repo: userRepo,
	}
}

func (a *AuthUseCase) SignIn(username, password string) {
	
}

func (a *AuthUseCase) SignUp(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}

	return a.repo.CreateUser(user)
}
