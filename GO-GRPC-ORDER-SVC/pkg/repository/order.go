package repository

import (
	"order/service/pkg/domain"
	"order/service/pkg/util/models"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{
		DB: DB,
	}
}

func (or *orderRepository) DoesCartExist(userID int) (bool, error) {

	var exist bool
	err := or.DB.Raw("select exists(select 1 from carts where user_id = ?)", userID).Scan(&exist).Error
	if err != nil {
		return false, err
	}

	return exist, nil
}
func (or *orderRepository) AddressExist(orderBody models.OrderIncoming) (bool, error) {

	var count int
	if err := or.DB.Raw("SELECT COUNT(*) FROM addresses WHERE user_id = ? AND id = ?", orderBody.UserID, orderBody.AddressID).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil

}
func (or *orderRepository) PaymentExist(orderBody models.OrderIncoming) (bool, error) {
	var count int
	if err := or.DB.Raw("SELECT count(*) FROM payment_methods WHERE id = ?", orderBody.PaymentID).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil

}
func (or *orderRepository) PaymentStatus(orderID int) (string, error) {
	var status string
	err := or.DB.Raw("SELECT payment_status FROM orders WHERE id= ?", orderID).Scan(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}
func (or *orderRepository) TotalAmountInCart(userID int) (float64, error) {
	var price float64
	if err := or.DB.Raw("SELECT sum(total_price) FROM carts WHERE  user_id= $1", userID).Scan(&price).Error; err != nil {
		return 0, err
	}
	return price, nil
}
func (or *orderRepository) OrderItems(ob models.OrderIncoming, price float64) (int, error) {
	var id int
	query := `
    INSERT INTO orders (created_at , user_id , address_id , payment_method_id , final_price)
    VALUES (NOW(),?, ?, ?, ?)
    RETURNING id`
	or.DB.Raw(query, ob.UserID, ob.AddressID, ob.PaymentID, price).Scan(&id)
	return id, nil
}
func (or *orderRepository) AddOrderProducts(order_id int, cart []models.Cart) error {
	query := `
    INSERT INTO order_items (order_id,product_id,quantity,total_price)
    VALUES (?, ?, ?, ?) `
	for _, v := range cart {
		var productID int
		if err := or.DB.Raw("SELECT id FROM products WHERE name = $1", v.ProductName).Scan(&productID).Error; err != nil {
			return err
		}
		if err := or.DB.Exec(query, order_id, productID, v.Quantity, v.TotalPrice).Error; err != nil {
			return err
		}
	}
	return nil
}
func (or *orderRepository) GetBriefOrderDetails(orderID int) (domain.OrderSuccessResponse, error) {
	var orderSuccessResponse domain.OrderSuccessResponse
	err := or.DB.Raw(`SELECT id as order_id,shipment_status FROM orders WHERE id = ?`, orderID).Scan(&orderSuccessResponse).Error
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	return orderSuccessResponse, nil
}
func (or *orderRepository) UpdateCartAfterOrder(userID, productID int, quantity float64) error {
	err := or.DB.Exec("DELETE FROM carts WHERE user_id = ? and product_id = ?", userID, productID).Error
	if err != nil {
		return err
	}

	err = or.DB.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", quantity, productID).Error
	if err != nil {
		return err
	}

	return nil
}
func (or *orderRepository) GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var orderDetails []models.OrderDetails
	or.DB.Raw("SELECT id as order_id,final_price,shipment_status,payment_status FROM orders WHERE user_id = ? LIMIT ? OFFSET ? ", userId, count, offset).Scan(&orderDetails)
	var fullOrderDetails []models.FullOrderDetails
	for _, od := range orderDetails {
		var orderProductDetails []models.OrderProductDetails
		or.DB.Raw(`SELECT
		order_items.product_id,
		products.name AS product_name,
		order_items.quantity,
		order_items.total_price
	    FROM
		order_items
	    INNER JOIN
		products ON order_items.product_id = products.id
	    WHERE
		order_items.order_id = $1 `, od.OrderId).Scan(&orderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.FullOrderDetails{OrderDetails: od, OrderProductDetails: orderProductDetails})
	}
	return fullOrderDetails, nil
}