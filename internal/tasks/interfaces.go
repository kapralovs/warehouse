package tasks

import (
	"github.com/kapralovs/warehouse/internal/products"
)

// Интерфейс задачи
type Task interface {
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
	ID            string // ID заказа
	Client        string // Клиент
	Destination   string // Пункт назначения
	Items         []*products.Product
	ShippingUnits []*ShippingUnit
}

// Структура задачи пополнения
type Refill struct {
	Header         *TaskHeader
	ReleasingPlace string
	ProductsID     []string
	Count          int
	RecievingPlace string
}

// Структура задачи отгрузки
type Transport struct {
	RecievingPlace string
}
