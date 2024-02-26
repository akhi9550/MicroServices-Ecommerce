package client

import (
	"context"
	"errors"
	"fmt"
	interfaces "grpc-api-gateway/pkg/client/interface"
	"grpc-api-gateway/pkg/config"
	pb "grpc-api-gateway/pkg/pb/cart"
	"grpc-api-gateway/pkg/utils/models"

	"google.golang.org/grpc"
)

type cartClient struct {
	Client pb.CartClient
}

func NewCartClient(cfg config.Config) interfaces.CartClient {

	grpcConnection, err := grpc.Dial(cfg.CartSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewCartClient(grpcConnection)

	return &cartClient{
		Client: grpcClient,
	}

}
func (c *cartClient) AddToCart(product_id int, user_id int) (models.CartResponse, error) {
	res, err := c.Client.AddToCart(context.Background(), &pb.AddToCartRequest{
		ProductID: int64(product_id),
		UserID:    int64(user_id),
	})
	if err != nil {
		return models.CartResponse{}, err
	}
	if res.Error != "" {
		return models.CartResponse{}, errors.New(res.Error)
	}
	cart := models.Cart{
		ProductID:   uint(res.Cart.ProductID),
		ProductName: res.Cart.ProductName,
		Quantity:    float64(res.Cart.Quantity),
		TotalPrice:  float64(res.Cart.TotalPrice),
	}
	return models.CartResponse{
		UserName:   res.Username,
		TotalPrice: float64(res.Price),
		Cart:       []models.Cart{cart},
	}, nil
}
func (c *cartClient) DisplayCart(user_id int) (models.CartResponse, error) {
	res, err := c.Client.GetCart(context.Background(), &pb.GetCartRequest{
		UserID: int64(user_id),
	})
	if err != nil {
		return models.CartResponse{}, err
	}
	if res.Error != "" {
		return models.CartResponse{}, errors.New(res.Error)
	}
	cart := models.Cart{
		ProductID:   uint(res.Cart.ProductID),
		ProductName: res.Cart.ProductName,
		Quantity:    float64(res.Cart.Quantity),
		TotalPrice:  float64(res.Cart.TotalPrice),
	}
	return models.CartResponse{
		UserName:   res.Username,
		TotalPrice: float64(res.Price),
		Cart:       []models.Cart{cart},
	}, nil
}
