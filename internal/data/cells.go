package data

import (
	"errors"
	"log"

	"github.com/kapralovs/warehouse/internal/models"
)

func (ds *DataStorage) SaveCell(c *models.Cell) error {
	if c != nil {
		if c.ID == "" {
			return errors.New("can't save cell with empty ID field")
		}

		ds.mu.Lock()
		defer ds.mu.Unlock()
		ds.Cells[c.ID] = c

		log.Printf("The cell \"%s\" is saved.\n", ds.Cells[c.ID].ID)
		return nil
	}

	return errors.New("can't save current cell, because profile is nil")
}

func (ds *DataStorage) LoadCell(id string) (*models.Cell, error) {
	if cell, ok := ds.Cells[id]; ok {
		return cell, nil
	}

	return nil, errors.New("cell with this id does not exist")
}

func (ds *DataStorage) DeleteCell(id string) error {
	if _, ok := ds.Cells[id]; ok {
		delete(ds.Cells, id)
		return nil
	}

	return errors.New("it is not possible to delete a cell because it does not exist")
}
