package tasks

import "github.com/kapralovs/warehouse/internal/users"

// Конструктор для нового заказа
func NewOrder() *Order {
	return &Order{}
}

func create() {
}

func delete() {

}

func edit() {

}

func AssignTo(usr *users.Profile) error {
	return nil
}

func ToPool() {}
