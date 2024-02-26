package service

import (
	pb "cart/service/pkg/pb/cart"
	interfaces "cart/service/pkg/usecase/interface"
	"context"
)

type CartServer struct {
	CartUseCase interfaces.CartUseCase
	pb.UnimplementedCartServer
}

func NewCartServer(useCase interfaces.CartUseCase) pb.CartServer {
	return &CartServer{
		CartUseCase: useCase,
	}
}

func (c *CartServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	productID := int(req.ProductID)
	userID := int(req.UserID)

	res, err := c.CartUseCase.AddToCart(productID, userID)
	if err != nil {
		return &pb.AddToCartResponse{
			Error: err.Error(),
		}, err
	}
	var result pb.AddToCartResponse
	for _, v := range res.Cart {
		var cartDetails pb.CartDetails
		cartDetails.ProductID = int64(v.ProductID)
		cartDetails.ProductName = v.ProductName
		cartDetails.Quantity = float32(v.Quantity)
		cartDetails.TotalPrice = float32(v.TotalPrice)

		response := &pb.AddToCartResponse{
			Username: res.UserName,
			Price:    float32(res.TotalPrice),
			Cart:     &cartDetails,
		}
		result.Username = response.Username
		result.Price = response.Price
		result.Cart = response.Cart
	}

	return &result, nil
}

func (c *CartServer) DisplayCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	userID := int(req.UserID)
	res, err := c.CartUseCase.DisplayCart(userID)
	if err != nil {
		return &pb.GetCartResponse{
			Error: err.Error(),
		}, err
	}
	var result pb.GetCartResponse
	for _, v := range res.Cart {
		var cartDetails pb.CartDetails
		cartDetails.ProductID = int64(v.ProductID)
		cartDetails.ProductName = v.ProductName
		cartDetails.Quantity = float32(v.Quantity)
		cartDetails.TotalPrice = float32(v.TotalPrice)

		response := &pb.GetCartResponse{
			Username: res.UserName,
			Price:    float32(res.TotalPrice),
			Cart:     &cartDetails,
		}
		result.Username = response.Username
		result.Price = response.Price
		result.Cart = response.Cart
	}

	return &result, nil
}
