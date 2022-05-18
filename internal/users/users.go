package users

import (
	"fmt"
	"log"
)

// Конструктор для нового пользователя
func New() *Profile {
	return &Profile{}
}

func CheckAdminRights(profile *Profile) error {
	if !profile.Account.IsAdmin {
		log.Printf("The administrator rights check failed. User \"%s\" is not an administrator\n", profile.Account.Username)
		return fmt.Errorf("user \"%s\" does not have administrator rights", profile.Account.Username)
	}

	log.Printf("The administrator rights check has been passed. User \"%s\" is administrator\n", profile.Account.Username)
	return nil
}
