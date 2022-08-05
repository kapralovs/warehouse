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

func New()*ProductRepository{
	return &ProductRepository{
		
	}
}

func (pr *ProductRepository) SaveProduct(p *models.Product) error {
	if p != nil {
		if p.ID == "" {
			return errors.New("can't save product with empty ID field")
		}

		pr.mu.Lock()
		defer pr.mu.Unlock()
		pr.Products[p.ID] = p

		log.Printf("The product \"%s\" is saved.\n", pr.Products[p.ID].ID)
		return nil
	}

	return errors.New("can't save current profile, because profile is nil")
}

func (pr *ProductRepository) LoadProduct(id string) (*models.Product, error) {
	if product, ok := pr.Products[id]; ok {
		return product, nil
	}

	return nil, errors.New("product with this id does not exist")
}

func (pr *ProductRepository) DeleteProduct(id string) error {
	if _, ok := pr.Products[id]; ok {
		delete(pr.Products, id)
		return nil
	}

	return errors.New("it is not possible to delete a product because it does not exist")
}
