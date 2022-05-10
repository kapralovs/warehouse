package warehouse

import "github.com/kapralovs/warehouse/internal/products"

// Структура склада
type Warehouse struct {
	ID    int
	Title string
	Cells map[string]*Cell
}

// Структура ячейки хранения продукта
type Cell struct {
	ID      string
	Content map[string]*products.Product
}
