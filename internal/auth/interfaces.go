package auth

import "github.com/kapralovs/warehouse/internal/models"

type AuthController interface {
	SignIn()
	SignUp()
}

type UseCase interface {
	SignIn()
	SignUp()
}

type Repository interface {
	CreateUser()
	GetUser(id string) (*models.User, error)
}
