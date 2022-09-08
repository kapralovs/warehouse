package usecase

import (
	"github.com/kapralovs/warehouse/internal/auth"
	"github.com/kapralovs/warehouse/internal/models"
)

type UseCase struct {
	repo auth.Repository
}

func New(userRepo auth.Repository) *UseCase {
	return &UseCase{
		repo: userRepo,
	}
}

func (a *UseCase) SignIn(username, password string) error {
	user, err := a.repo.GetUser(username, password)
	if err != nil {
		return err
	}

	// if profile, ok := a.Profiles[username]; ok {

	// 	if profile.Account.Password == password {
	// 		ds.Profiles[username].Account.IsOnline = true
	// 		fmt.Println("Profile pointer", profile)                             //DEBUG
	// 		fmt.Println("ds.Profiles[username] pointer", ds.Profiles[username]) //DEBUG
	// 		log.Printf("User \"%s\" is authorised\n", profile.Account.Username)
	// 		return profile, nil
	// 	}
	// }
	// if user, ok := checkCredentials(ds, username, password); ok {
	// 	log.Printf("User \"%s\" is authorised\n", user.Account.Username)
	// 	return user, nil
	// }
}

func (a *UseCase) SignUp(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}

	return a.repo.CreateUser(user)
}
