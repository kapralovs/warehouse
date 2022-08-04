package repository

import (
	"errors"
	"log"
	"sync"

	"github.com/kapralovs/warehouse/internal/models"
)

type ProductRepository struct {
	mu       *sync.Mutex
	Products map[string]*models.Product
}

func (ds *DataStorage) SaveProduct(p *models.Product) error {
	if p != nil {
		if p.ID == "" {
			return errors.New("can't save product with empty ID field")
		}

		ds.mu.Lock()
		defer ds.mu.Unlock()
		ds.Products[p.ID] = p

		log.Printf("The product \"%s\" is saved.\n", ds.Products[p.ID].ID)
		return nil
	}

	return errors.New("can't save current profile, because profile is nil")
}

func (ds *DataStorage) LoadProduct(id string) (*models.Product, error) {
	if product, ok := ds.Products[id]; ok {
		return product, nil
	}

	return nil, errors.New("product with this id does not exist")
}

func (ds *DataStorage) DeleteProduct(id string) error {
	if _, ok := ds.Products[id]; ok {
		delete(ds.Products, id)
		return nil
	}

	return errors.New("it is not possible to delete a product because it does not exist")
}
