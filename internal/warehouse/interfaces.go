package warehouse

import "github.com/kapralovs/warehouse/internal/products"

type Warehouse struct {
	ID    int
	Title string
	Cells map[string]*Cell
}

type Cell struct {
	ID      string
	Content []*products.Product
}
