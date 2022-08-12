package usecase

import (
	"github.com/kapralovs/warehouse/internal/models"
	"github.com/kapralovs/warehouse/internal/warehouse"
)

type WarehouseUseCase struct {
	warehouseRepo warehouse.Repository
}

// Функция проверяет содержимое ячейки
func CheckCellContent() []*models.Product {
	var content []*models.Product

	return content
}
