package interfaces

import "grpc-api-gateway/pkg/utils/models"

type CartClient interface {
	AddToCart(product_id int, user_id int) (models.CartResponse, error)
	DisplayCart(user_id int) (models.CartResponse, error)
}
