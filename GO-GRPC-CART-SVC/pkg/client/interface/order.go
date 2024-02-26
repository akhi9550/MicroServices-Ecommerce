package interfaces

type NewOrderClient interface {
	ProductStockMinus(productID, stock int) error
}
