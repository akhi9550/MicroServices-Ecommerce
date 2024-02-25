package interfaces
type OrderClient interface{
	ProductStockMinus(productID, stock int) error

}