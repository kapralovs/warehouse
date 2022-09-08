package usecase

import "github.com/kapralovs/warehouse/internal/products"

type UseCase struct {
	productRepo products.Repository
}

func New(repo products.Repository) *UseCase {
	return &UseCase{
		productRepo: repo,
	}
}

// Функция проверяет остаток на месте, чтобы при пополнении и/или заказе забрать остаток
// и освободить место для новых поступлений продукта
func (p *UseCase) CheckProductReserve(prodID string) map[string]int {
	productReserve := map[string]int{}
	for id, cell := range ds.Cells {
		if product, ok := cell.Content[prodID]; ok {
			productReserve[id] = product.Count
		}
	}
	return productReserve
}
