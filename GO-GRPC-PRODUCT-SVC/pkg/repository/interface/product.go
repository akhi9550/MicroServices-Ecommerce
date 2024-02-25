package interfaces

import (
	"product/service/pkg/domain"
	"product/service/pkg/utils/models"
)

type ProductRepository interface {
	ShowAllProducts(page int, count int) ([]models.ProductBrief, error)
	AddProducts(product models.Product) (domain.Product, error)
	ProductAlreadyExist(Name string) bool
	StockInvalid(Name string) bool
	DeleteProducts(id int) error
	CheckProductExist(pid int) (bool, error)
	UpdateProduct(pid int, stock int) (models.ProductUpdateReciever, error)
}
