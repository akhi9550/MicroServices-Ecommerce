package client

import (
	"cart/service/pkg/config"
	pb "cart/service/pkg/pb/order"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type clientOrder struct {
	Client pb.OrderClient
}

func NewOrderClient(c *config.Config) *clientOrder {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	pbClient := pb.NewOrderClient(cc)

	return &clientOrder{
		Client: pbClient,
	}
}
func (c *clientOrder) ProductStockMinus(productID, stock int) error {
	_, err := c.Client.ProductStockMinus(context.Background(), &pb.ProductStockMinusRequest{
		ID:    int64(productID),
		Stock: int64(stock),
	})
	if err != nil {
		return err
	}
	return nil
}
