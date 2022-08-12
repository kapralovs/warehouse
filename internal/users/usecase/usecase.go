package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/kapralovs/warehouse/internal/models"
	"github.com/kapralovs/warehouse/internal/users"
)

type UserUseCase struct {
	userRepo users.Repository
}

func New(repo users.Repository) *UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

// Проверка на наличие прав администратора
func CheckAdminRights(user *models.User) error {
	if !user.IsAdmin {
		log.Printf("The administrator rights check failed. User \"%s\" is not an administrator\n", user.Username)
		return fmt.Errorf("user \"%s\" does not have administrator rights", user.Username)
	}

	log.Printf("The administrator rights check has been passed. User \"%s\" is administrator\n", user.Username)
	return nil
}

// Проверка на online-статус
func CheckOnlineStatus(p *models.User) error {
	if p.IsOnline {
		return nil
	}

	return errors.New("this user is offline")
}
