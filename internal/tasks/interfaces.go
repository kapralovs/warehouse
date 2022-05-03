package tasks

// Интерфейс задачи
type Task interface {
	Start()
	Complete()
}

// Структуры, определяющие конкретные задачи

// Структура задачи заказа клиента
type Order struct {
}

// Структура задачи пополнения
type Refill struct {
}

// Структура задачи отгрузки
type Shipment struct {
}
