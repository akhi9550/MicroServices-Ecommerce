package interfaces

import "order/service/pkg/util/models"

type CartClient interface {
	GetAllItemsFromCart(userID int) ([]models.Cart, error)
}
