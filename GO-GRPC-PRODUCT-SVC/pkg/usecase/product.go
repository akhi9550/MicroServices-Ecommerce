package usecase

import (
	"errors"
	"product/service/pkg/domain"
	interfaces "product/service/pkg/repository/interface"
	services "product/service/pkg/usecase/interface"
	"product/service/pkg/utils/models"
)

type productUseCase struct {
	productRepository interfaces.ProductRepository
}

func NewProductUseCase(repository interfaces.ProductRepository) services.ProductUseCase {
	return &productUseCase{
		productRepository: repository,
	}
}
func (pr *productUseCase) ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {
	productDetails, err := pr.productRepository.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductBrief{}, err
	}

	for i := range productDetails {
		p := &productDetails[i]
		if p.Stock <= 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}

	return productDetails, nil
}

func (pr *productUseCase) AddProducts(product models.Product) (domain.Product, error) {
	exist := pr.productRepository.ProductAlreadyExist(product.Name)
	if exist {
		return domain.Product{}, errors.New("product already exist")
	}
	productResponse, err := pr.productRepository.AddProducts(product)
	if err != nil {
		return domain.Product{}, err
	}
	stock := pr.productRepository.StockInvalid(productResponse.Name)
	if !stock {
		return domain.Product{}, errors.New("stock is invalid input")
	}
	return productResponse, nil
}
func (pr *productUseCase) DeleteProducts(id int) error {
	err := pr.productRepository.DeleteProducts(id)
	if err != nil {
		return err
	}
	return nil
}
func (pr *productUseCase) UpdateProduct(pid int, stock int) (models.ProductUpdateReciever, error) {
	if stock <= 0 {
		return models.ProductUpdateReciever{}, errors.New("stock doesnot update invalid input")
	}
	result, err := pr.productRepository.CheckProductExist(pid)
	if err != nil {
		return models.ProductUpdateReciever{}, err
	}
	if !result {
		return models.ProductUpdateReciever{}, errors.New("there is no product as you mentioned")
	}
	newcat, err := pr.productRepository.UpdateProduct(pid, stock)
	if err != nil {
		return models.ProductUpdateReciever{}, err
	}
	return newcat, err

}
