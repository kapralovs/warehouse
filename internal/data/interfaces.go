package data

import (
	"sync"

	"github.com/kapralovs/warehouse/internal/models"
)

// type Saver interface {
// 	Save()
// }

// type Loader interface {
// 	Load()
// }

// type Creator interface {
// 	Create()
// }

// Структура in-memory базы данных
type DataStorage struct {
	mu         sync.Mutex
	Tasks      map[string]models.Task
	Products   map[string]*models.Product
	Profiles   map[string]*models.Profile
	Cells      map[string]*models.Cell
	Warehouses map[string]*models.Warehouse
}
