package repository

import (
	"errors"
	"log"
	"sync"

	"github.com/kapralovs/warehouse/internal/models"
)

type WarehouseRepository struct {
	mu         *sync.Mutex
	Cells      map[string]*models.Cell
	Warehouses map[string]*models.Warehouse
}

func (wr *WarehouseRepository) SaveCell(c *models.Cell) error {
	if c != nil {
		if c.ID == "" {
			return errors.New("can't save cell with empty ID field")
		}

		wr.mu.Lock()
		defer wr.mu.Unlock()
		wr.Cells[c.ID] = c

		log.Printf("The cell \"%s\" is saved.\n", wr.Cells[c.ID].ID)
		return nil
	}

	return errors.New("can't save current cell, because profile is nil")
}

func (wr *WarehouseRepository) LoadCell(id string) (*models.Cell, error) {
	if cell, ok := wr.Cells[id]; ok {
		return cell, nil
	}

	return nil, errors.New("cell with this id does not exist")
}

func (wr *WarehouseRepository) DeleteCell(id string) error {
	if _, ok := wr.Cells[id]; ok {
		delete(wr.Cells, id)
		return nil
	}

	return errors.New("it is not possible to delete a cell because it does not exist")
}
