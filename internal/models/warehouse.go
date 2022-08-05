package models

// Структура склада
type Warehouse struct {
	// ID склада
	ID int
	// Наименование склада
	Title string
	// Список всех мест склада
	Cells map[string]*Cell
}

// Структура ячейки хранения продукта
type Cell struct {
	// ID ячейки
	ID string
	// Содержимое ячейки
	Content map[string]*Product
}
