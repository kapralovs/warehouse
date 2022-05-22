package data

import (
	"errors"
	"log"

	"github.com/kapralovs/warehouse/internal/products"
	"github.com/kapralovs/warehouse/internal/users"
)

// Конструкор для нового хранилища данных
func NewStorage() *DataStorage {
	return &DataStorage{}
}

// Инициализация хранилища
func (ds *DataStorage) Init() {
	profilesDB := make(map[string]*users.Profile, 3)
	ds.Profiles = profilesDB
	ds.Profiles["1"] = &users.Profile{
		General: &users.GeneralInfo{
			Lastname: "Adminov",
			Name:     "Admin",
			Age:      36,
			Position: users.ManagerPosition,
			Salary:   50000,
		},
		Account: &users.Account{
			ID:       "1",
			Username: "admin",
			Password: "admin",
			IsAdmin:  true,
		},
		Productivity: &users.Productivity{
			PositionsCounter: 234,
		},
	}
}

func CheckDuplicates() error {
	return nil
}

// Функция проверяет остаток на месте, чтобы при пополнении и/или заказе забрать остаток
// и освободить место для новых поступлений продукта
func (ds *DataStorage) CheckProductReserve(prodID string) map[string]int {
	productReserve := map[string]int{}
	for id, cell := range ds.Cells {
		if product, ok := cell.Content[prodID]; ok {
			productReserve[id] = product.Count
		}
	}
	return productReserve
}

// Функция проверяет содержимое ячейки
func CheckCellContent() []*products.Product {
	var content []*products.Product

	return content
}

func (ds *DataStorage) CheckProfileDuplicates(p *users.Profile) error {
	if p != nil {
		if p.Account.ID != "" {
			for id, profile := range ds.Profiles {
				if id == p.Account.ID {
					return errors.New("profile with this ID already exists")
				}
				if profile.Account.Username == p.Account.Username {
					return errors.New("profile with this username already exists")
				}

				return nil
			}

			log.Println()
			return errors.New("empty profile ID")
		}
	}

	return errors.New("edited user profile is nil")
}
