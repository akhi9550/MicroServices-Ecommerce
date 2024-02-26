package client

import (
	"cart/service/pkg/config"
	pb "cart/service/pkg/pb/product"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type clientProduct struct {
	Client pb.ProductClient
}

func NewProductClient(c *config.Config) *clientProduct {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	pbClient := pb.NewProductClient(cc)

	return &clientProduct{
		Client: pbClient,
	}
}
func (c *clientProduct) GetQuantityFromProductID(id int) (int, error) {
	res, err := c.Client.GetQuantityFromProductID(context.Background(), &pb.GetQuantityFromProductIDRequest{
		ID: int64(id),
	})
	if err != nil {
		return 0, err
	}
	quantity := int(res.Quantity)
	return quantity, nil
}
func (c *clientProduct) GetPriceOfProductFromID(prodcut_id int) (float64, error) {
	res, err := c.Client.GetPriceofProductFromID(context.Background(), &pb.GetPriceofProductFromIDRequest{
		ID: int64(prodcut_id),
	})
	if err != nil {
		return 0, err
	}
	price := float64(res.Price)
	return price, nil
}
