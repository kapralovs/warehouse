package usecase

import "github.com/kapralovs/warehouse/internal/auth"

type AuthUseCase struct {
	Repository auth.Repository
}
