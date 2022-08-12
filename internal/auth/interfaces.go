package auth

import "github.com/kapralovs/warehouse/internal/models"

type AuthController interface {
	SignIn()
	SignUp()
}

type AuthUseCase interface {
	SignIn(username, password string) error
	SignUp(username, password string) (error, string)
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(username, password string) (*models.User, error)
}
