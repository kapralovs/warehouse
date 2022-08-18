package auth

import "github.com/kapralovs/warehouse/internal/models"

type AuthController interface {
	SignIn()
	SignUp()
}

type UseCase interface {
	SignIn(username, password string) error
	SignUp(username, password string) (error, string)
}

type Repository interface {
	CreateUser(user *models.User) error
	GetUser(username, password string) (*models.User, error)
}
