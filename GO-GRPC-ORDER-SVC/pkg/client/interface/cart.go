package interfaces

type CartRepository interface {
	GetAllItemsFromCart(userID int) ([]models.Cart, error)
}
