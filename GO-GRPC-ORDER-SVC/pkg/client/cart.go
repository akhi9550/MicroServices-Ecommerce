package client

import (
	"errors"
	"order/service/pkg/util/models"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: DB,
	}
}

func (cr *cartRepository) GetAllItemsFromCart(userID int) ([]models.Cart, error) {
	var count int
	var cartResponse []models.Cart
	err := cr.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return []models.Cart{}, err
	}
	if count == 0 {
		return []models.Cart{}, nil
	}
	err = cr.DB.Raw("SELECT carts.user_id,users.firstname as user_name,carts.product_id,products.name as product_name,carts.quantity,carts.total_price from carts INNER JOIN users on carts.user_id = users.id INNER JOIN products ON carts.product_id = products.id where user_id = ?", userID).First(&cartResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if len(cartResponse) == 0 {
				return []models.Cart{}, nil
			}
			return []models.Cart{}, err
		}
		return []models.Cart{}, err
	}
	return cartResponse, nil
}
