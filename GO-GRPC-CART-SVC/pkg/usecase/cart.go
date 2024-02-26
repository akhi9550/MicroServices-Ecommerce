package usecase

import (
	clinetinterface "cart/service/pkg/client/interface"
	interfaces "cart/service/pkg/repository/interface"
	services "cart/service/pkg/usecase/interface"
	"cart/service/pkg/utils/models"
	"errors"
)

type cartUseCase struct {
	cartRepository    interfaces.CartRepository
	productRepository clinetinterface.NewProductClient
	orderRepository   clinetinterface.NewOrderClient
}

func NewCartUseCase(repository interfaces.CartRepository, productRepo clinetinterface.NewProductClient, orderRepo clinetinterface.NewOrderClient) services.CartUseCase {

	return &cartUseCase{
		cartRepository:    repository,
		productRepository: productRepo,
		orderRepository:   orderRepo,
	}

}
func (cr *cartUseCase) AddToCart(product_id int, user_id int) (models.CartResponse, error) {
	ok, _, err := cr.cartRepository.CheckProduct(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, errors.New("product Does not exist")
	}
	QuantityOfProductInCart, err := cr.cartRepository.QuantityOfProductInCart(user_id, product_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	quantityOfProduct, err := cr.productRepository.GetQuantityFromProductID(product_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	if quantityOfProduct <= 0 {
		return models.CartResponse{}, errors.New("out of stock")
	}
	if quantityOfProduct == QuantityOfProductInCart {
		return models.CartResponse{}, errors.New("stock limit exceeded")
	}
	productPrice, err := cr.productRepository.GetPriceOfProductFromID(product_id)
	if err != nil {

		return models.CartResponse{}, err
	}

	FinalPrice := productPrice
	if QuantityOfProductInCart == 0 {
		err := cr.cartRepository.AddItemIntoCart(user_id, product_id, 1, FinalPrice)
		if err != nil {

			return models.CartResponse{}, err
		}

	} else {
		currentTotal, err := cr.cartRepository.TotalPriceForProductInCart(user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
		err = cr.cartRepository.UpdateCart(QuantityOfProductInCart+1, currentTotal+productPrice, user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}
	cartDetails, err := cr.cartRepository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := cr.cartRepository.GetTotalPrice(user_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	err = cr.orderRepository.ProductStockMinus(product_id, QuantityOfProductInCart)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cartDetails,
	}, nil

}
func (cr *cartUseCase) DisplayCart(user_id int) (models.CartResponse, error) {
	cart, err := cr.cartRepository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := cr.cartRepository.GetTotalPrice(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cart,
	}, nil
}
