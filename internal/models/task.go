package models

// Интерфейс задачи
type Task interface {
	ID() string
	Start() error
	Complete() error
}

// Структура заголовка задачи
type TaskHeader struct {
	// Код склада, определяющий куда будет происходить пополнение (на пик-зону текущего склада или на удаленный склад)
	WarehouseID string
	// ID задачи в системе
	ID string
	// Название задачи
	Title string
	// Кем был создан заказ (ID менеджера, диспетчера, либо другого сотрудника, уполномоченного создавать новые задачи).
	CreatedBy string
	// Исполнитель заказа
	Resource string
}

// Структура Единицы Отгрузки (ЕО), которая хранит информацию о товарах, которые были закреплены
type ShippingUnit struct {
	Items []*Product
}

// --------------------Структуры, определяющие конкретные задачи--------------------

// Структура задачи заказа клиента
type Order struct {
	Header *TaskHeader
	// ID заказа
	ID string `json:"id"`
	// Клиент
	Client string `json:"client"`
	// Пункт назначения
	Destination string `json:"destination"`
	// Позиции заказа
	Items []*Product
	// Единицы отгрузки
	ShippingUnits []*ShippingUnit
}

// Структура задачи пополнения
type Refill struct {
	Header *TaskHeader
	// Отпускающее место
	ReleasingPlace string `json:"releasing_place"`
	// ID требуемого продукта
	ProductID string `json:"product_id"`
	// Требуемое кол-во продукта
	Count int `json:"count"`
	//Принимающее складское место
	RecievingPlace string `json:"recieving_place"`
}

// Структура задачи отгрузки
type Transport struct {
	// Место хранения готового заказа
	RecievingPlace string `json:"recieving_place"`
}
