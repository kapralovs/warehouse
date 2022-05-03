package warehouse

import "github.com/kapralovs/warehouse/internal/products"

type Warehouse struct {
	ID string
}

type Cell struct {
	ID      string
	Stage   string
	Content []*products.Product
}
