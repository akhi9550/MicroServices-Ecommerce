package service

import (
	"context"
	"product/service/pkg/pb"
	interfaces "product/service/pkg/usecase/interface"
	"product/service/pkg/utils/models"
)

type ProductServer struct {
	productUseCase interfaces.ProductUseCase
	pb.UnimplementedProductServer
}

func NewProductServer(useCase interfaces.ProductUseCase) pb.ProductServer {
	return &ProductServer{
		productUseCase: useCase,
	}
}
func (p *ProductServer) AddProduct(ctx context.Context, details *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	model := models.Product{
		Name:        details.Name,
		Description: details.Description,
		CategoryID:  uint(details.CategoryID),
		Size:        int(details.Size),
		Stock:       int(details.Stock),
		Price:       float64(details.Price),
	}
	data, err := p.productUseCase.AddProducts(model)
	if err != nil {
		return &pb.AddProductResponse{}, err
	}
	return &pb.AddProductResponse{
		ID:          int64(data.ID),
		Name:        data.Name,
		Description: data.Description,
		CategoryID:  int64(data.CategoryID),
		Size:        int64(data.Size),
		Stock:       int64(data.Stock),
		Price:       float32(data.Price),
	}, nil
}
func (p *ProductServer) ListProducts(ctx context.Context, details *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	page := int(details.Page)
	count := int(details.Count)

	products, err := p.productUseCase.ShowAllProducts(page, count)
	if err != nil {
		return &pb.ListProductResponse{
			Details: []*pb.ProductDetails{},
		}, err
	}
	var result pb.ListProductResponse
	for _, v := range products {
		result.Details = append(result.Details, &pb.ProductDetails{
			ID:          int64(v.ID),
			Name:        v.Name,
			Description: v.Description,
			CategoryID:  int64(v.CategoryID),
			Size:        int64(v.Size),
			Stock:       int64(v.Stock),
			Price:       float32(v.Price),
		})
	}
	return &result, nil
}
func (p *ProductServer) UpdateProduct(ctx context.Context, details *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	id := int(details.ID)
	stock := int(details.Stock)
	update, err := p.productUseCase.UpdateProduct(id, stock)
	if err != nil {
		return &pb.UpdateProductResponse{}, err
	}
	return &pb.UpdateProductResponse{
		ID:    int64(update.ProductID),
		Stock: int64(update.Stock),
	}, nil
}
func (p *ProductServer) DeleteProduct(ctx context.Context, details *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	id := int(details.ID)
	err := p.productUseCase.DeleteProducts(id)
	if err != nil {
		return &pb.DeleteProductResponse{
			Error: err.Error(),
		}, err
	}
	return &pb.DeleteProductResponse{}, nil
}
