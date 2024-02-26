package handler

import (
	interfaces "grpc-api-gateway/pkg/client/interface"
	"grpc-api-gateway/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type CartHandler struct {
	GRPC_Client interfaces.CartClient
}

func NewCartHandler(client interfaces.CartClient) *CartHandler {
	return &CartHandler{
		GRPC_Client: client,
	}
}
func (ct *CartHandler) AddToCart(c *gin.Context) {
	id := c.Query("product_id")
	product_id, err := strconv.Atoi(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadGateway, "Product id is given in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	user_ID, _ := c.Get("user_id")
	cartResponse, err := ct.GRPC_Client.AddToCart(product_id, user_ID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadGateway, "could not add product to the cart", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Added porduct Successfully to the cart", cartResponse, nil)
	c.JSON(http.StatusOK, success)
}
func (ct *CartHandler) GetCart(c *gin.Context) {
	userID, _ := c.Get("user_id")
	cart, err := ct.GRPC_Client.DisplayCart(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "cannot display cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Cart items displayed successfully", cart, nil)
	c.JSON(http.StatusOK, success)
}
