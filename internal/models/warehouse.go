package models

// Структура склада
type Warehouse struct {
	ID    int              // ID склада
	Title string           //Наименование склада
	Cells map[string]*Cell // Список всех мест склада
}

// Структура ячейки хранения продукта
type Cell struct {
	ID      string              // ID ячейки
	Content map[string]*Product // Содержимое ячейки
}
