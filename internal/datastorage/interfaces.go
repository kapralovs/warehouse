package storage

import (
	"sync"

	"github.com/kapralovs/warehouse/internal/products"
	"github.com/kapralovs/warehouse/internal/tasks"
)

// Структура in memory-базы данных
type DataStorage struct {
	mu       sync.Mutex
	tasks    map[string]tasks.Task
	products map[string]*products.Product
}
