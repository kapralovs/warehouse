package usecase

import "github.com/kapralovs/warehouse/internal/warehouse"

type WarehouseUseCase struct {
	warehouseRepo warehouse.Repository
}
