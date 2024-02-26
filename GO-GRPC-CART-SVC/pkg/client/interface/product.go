package interfaces

type NewProductClient interface {
	GetQuantityFromProductID(id int) (int, error)
	GetPriceOfProductFromID(prodcut_id int) (float64, error)
}
