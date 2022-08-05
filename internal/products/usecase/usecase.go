package usecase

import "github.com/kapralovs/warehouse/internal/products"

type ProductUseCase struct {
	productRepo products.Repository
}
