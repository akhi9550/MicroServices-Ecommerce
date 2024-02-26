package client

import (
	"context"
	"fmt"
	clienterface "order/service/pkg/client/interface"
	"order/service/pkg/config"
	pb "order/service/pkg/pb/cart"
	"order/service/pkg/util/models"

	"google.golang.org/grpc"
)

type cartClient struct {
	Client pb.CartClient
}

func NewCartClient(cfg config.Config) clienterface.CartClient {

	grpcConnection, err := grpc.Dial(cfg.CartSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewCartClient(grpcConnection)

	return &cartClient{
		Client: grpcClient,
	}

}

func (c *cartClient) GetAllItemsFromCart(userID int) ([]models.Cart, error) {
	res, err := c.Client.GetCart(context.Background(), &pb.GetCartRequest{
		UserID: int64(userID),
	})
	if err != nil {
		return []models.Cart{}, err
	}
	var result []models.Cart

	for _, v := range res.Cart {
		result = append(result, models.Cart{
			ProductID:   uint(v.ProductID),
			ProductName: v.ProductName,
			Quantity:    float64(v.Quantity),
			TotalPrice:  float64(v.TotalPrice),
		})
	}
	fmt.Println(result)
	return result, nil
}
