package data

import (
	"fmt"

	"github.com/kapralovs/warehouse/internal/products"
)

func Create() error {
	return nil
}

func (ds *DataStorage) Load(id string) (*products.Product, error) {
	if product, ok := ds.Products[id]; ok {
		return product, nil
	}

	return nil, fmt.Errorf("")
}

func Delete() error {
	return nil
}
