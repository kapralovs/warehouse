package tasks

// Интерфейс задачи
type Task interface {
	Start()
	Complete()
}

// Структуры, определяющие конкретные задачи

// Структура задачи заказа клиента
type Order struct {
	ID string
}

// Структура задачи пополнения
type Refill struct {
}

// Структура задачи отгрузки
type Shipment struct {
}
