package tasks

import (
	"github.com/kapralovs/warehouse/internal/products"
)

// Интерфейс задачи
type Task interface {
	ID() string
	Start()
	Complete()
}

// Структура заголовка задачи
type TaskHeader struct {
	WarehouseID string // Код склада, определяющий куда будет происходить пополнение (на пик-зону текущего склада или на удаленный склад)
	ID          string // ID задачи в системе
	Title       string // Название задачи
	CreatedBy   string // Кем был создан заказ (ID менеджера, диспетчера, либо другого сотрудника, уполномоченного создавать новые задачи).
	Resource    string // исполнитель заказа
}

// Структура Единицы Отгрузки (ЕО), которая хранит информацию о товарах, которые были закреплены
type ShippingUnit struct {
	Items []*products.Product
}

// --------------------Структуры, определяющие конкретные задачи--------------------

// Структура задачи заказа клиента
type Order struct {
	Header        *TaskHeader
	ID            string `json:"id"`          // ID заказа
	Client        string `json:"client"`      // Клиент
	Destination   string `json:"destination"` // Пункт назначения
	Items         []*products.Product
	ShippingUnits []*ShippingUnit
}

// Структура задачи пополнения
type Refill struct {
	Header         *TaskHeader
	ReleasingPlace string `json:"releasing_place"` // Отпускающее место
	ProductID      string `json:"product_id"`      // ID требуемого продукта
	Count          int    `json:"count"`           // Требуемое кол-во продукта
	RecievingPlace string `json:"recieving_place"` //Принимающее складское место
}

// Структура задачи отгрузки
type Transport struct {
	RecievingPlace string `json:"recieving_place"` // Место хранения готового заказа
}
