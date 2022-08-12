package repository

import (
	"log"
	"sync"

	"github.com/kapralovs/warehouse/internal/models"
	"github.com/kapralovs/warehouse/internal/warehouse"
)

type WarehouseRepository struct {
	mu         *sync.Mutex
	Cells      map[string]*models.Cell
	Warehouses map[string]*models.Warehouse
}

func (wr *WarehouseRepository) SaveCell(c *models.Cell) error {
	if c != nil {
		if c.ID == "" {
			return warehouse.ErrEmptyId
		}

		wr.mu.Lock()
		defer wr.mu.Unlock()
		wr.Cells[c.ID] = c

		log.Printf("The cell \"%s\" is saved.\n", wr.Cells[c.ID].ID)
		return nil
	}

	return warehouse.ErrNilCell
}

func (wr *WarehouseRepository) GetCell(id string) (*models.Cell, error) {
	if cell, ok := wr.Cells[id]; ok {
		return cell, nil
	}

	return nil, warehouse.ErrCellNotExists
}

func (wr *WarehouseRepository) DeleteCell(id string) error {
	if _, ok := wr.Cells[id]; ok {
		delete(wr.Cells, id)
		return nil
	}

	return warehouse.ErrCellNotExists
}
