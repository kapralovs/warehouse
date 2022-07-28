package users

import (
	"errors"
	"fmt"
	"log"

	"github.com/kapralovs/warehouse/internal/models"
)

// Конструктор для нового пользователя
func New() *models.Profile {
	return &models.Profile{}
}

// Проверка на наличие прав администратора
func CheckAdminRights(profile *models.Profile) error {
	if !profile.Account.IsAdmin {
		log.Printf("The administrator rights check failed. User \"%s\" is not an administrator\n", profile.Account.Username)
		return fmt.Errorf("user \"%s\" does not have administrator rights", profile.Account.Username)
	}

	log.Printf("The administrator rights check has been passed. User \"%s\" is administrator\n", profile.Account.Username)
	return nil
}

// Проверка на online-статус
func CheckOnlineStatus(p *models.Profile) error {
	if p.Account.IsOnline {
		return nil
	}

	return errors.New("this user is offline")
}
