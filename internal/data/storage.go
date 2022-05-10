package data

import (
	"github.com/kapralovs/warehouse/internal/products"
	"github.com/kapralovs/warehouse/internal/users"
)

// Конструкор для нового хранилища данных
func New() *DataStorage {
	return &DataStorage{}
}

// Инициализация хранилища
func (ds *DataStorage) Init() {
	ds.Users["1"] = &users.User{
		General: &users.GeneralInfo{
			Lastname: "Adminov",
			Name:     "Admin",
			Age:      36,
			Position: users.ManagerPosition,
			Salary:   50000,
		},
		Account: &users.Account{
			ID:    "1",
			Login: "admin",
		},
	}
}

func CheckDuplicates() error {
	return nil
}

// Функция проверяет остаток на месте, чтобы при пополнении и/или заказе забрать остаток
// и освободить место для новых поступлений продукта
func (ds *DataStorage) CheckProductReserve(prodID string) map[string]int {
	for id, cell := range ds.Cells {
		if _, ok := cell.Content[prodID]; ok {

		}
	}
	productReserve := map[string]int{}
	return productReserve
}

func CheckCellContent() []*products.Product {
	var content []*products.Product

	return content
}
