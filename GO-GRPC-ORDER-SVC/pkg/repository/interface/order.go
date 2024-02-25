package interfaces

import (
	"order/service/pkg/domain"
	"order/service/pkg/util/models"
)

type OrderRepository interface {
	DoesCartExist(userID int) (bool, error)
	AddressExist(orderBody models.OrderIncoming) (bool, error)
	PaymentExist(orderBody models.OrderIncoming) (bool, error)
	PaymentStatus(orderID int) (string, error)
	TotalAmountInCart(userID int) (float64, error)
	OrderItems(ob models.OrderIncoming, price float64) (int, error)
	AddOrderProducts(order_id int, cart []models.Cart) error
	GetBriefOrderDetails(orderID int) (domain.OrderSuccessResponse, error)
	UpdateCartAfterOrder(userID, productID int, quantity float64) error
	GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error)
}
