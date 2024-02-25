package interfaces

import "cart/service/pkg/utils/models"

type CartRepository interface {
	DisplayCart(userID int) ([]models.Cart, error)
	GetTotalPrice(userID int) (models.CartTotal, error)
	CartExist(userID int) (bool, error)

	EmptyCart(userID int) error
	CheckProduct(product_id int) (bool, string, error)
	QuantityOfProductInCart(userId int, productId int) (int, error)
	AddItemIntoCart(userId int, productId int, Quantity int, productprice float64) error
	TotalPriceForProductInCart(userID int, productID int) (float64, error)
	UpdateCart(quantity int, price float64, userID int, product_id int) error
	ProductExist(userID int, productID int) (bool, error)
	GetAllItemsFromCart(userID int) ([]models.Cart, error)

}
