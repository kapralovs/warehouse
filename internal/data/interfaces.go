package data

import (
	"sync"

	"github.com/kapralovs/warehouse/internal/products"
	"github.com/kapralovs/warehouse/internal/tasks"
	"github.com/kapralovs/warehouse/internal/users"
	"github.com/kapralovs/warehouse/internal/warehouse"
)

// Структура in-memory базы данных
type DataStorage struct {
	mu         sync.Mutex
	Tasks      map[string]tasks.Task
	Products   map[string]*products.Product
	Users      map[string]*users.User
	Cells      map[string]*warehouse.Cell
	Warehouses map[string]*warehouse.Warehouse
}
