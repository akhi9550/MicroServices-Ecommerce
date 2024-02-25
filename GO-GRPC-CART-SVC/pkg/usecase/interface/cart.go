package interfaces

import "cart/service/pkg/utils/models"

type CartUseCase interface {
	AddToCart(product_id int, user_id int) (models.CartResponse, error)
	DisplayCart(user_id int) (models.CartResponse, error)
}
