// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: pkg/pb/product/product.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Product_AddProduct_FullMethodName               = "/product.Product/AddProduct"
	Product_ListProducts_FullMethodName             = "/product.Product/ListProducts"
	Product_UpdateProducts_FullMethodName           = "/product.Product/UpdateProducts"
	Product_DeleteProduct_FullMethodName            = "/product.Product/DeleteProduct"
	Product_GetQuantityFromProductID_FullMethodName = "/product.Product/GetQuantityFromProductID"
	Product_GetPriceofProductFromID_FullMethodName  = "/product.Product/GetPriceofProductFromID"
	Product_ProductStockMinus_FullMethodName        = "/product.Product/ProductStockMinus"
	Product_CheckProduct_FullMethodName             = "/product.Product/CheckProduct"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error)
	ListProducts(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error)
	UpdateProducts(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	GetQuantityFromProductID(ctx context.Context, in *GetQuantityFromProductIDRequest, opts ...grpc.CallOption) (*GetQuantityFromProductIDResponse, error)
	GetPriceofProductFromID(ctx context.Context, in *GetPriceofProductFromIDRequest, opts ...grpc.CallOption) (*GetPriceofProductFromIDResponse, error)
	ProductStockMinus(ctx context.Context, in *ProductStockMinusRequest, opts ...grpc.CallOption) (*ProductStockMinusReponse, error)
	CheckProduct(ctx context.Context, in *CheckProductRequest, opts ...grpc.CallOption) (*CheckProductResponse, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error) {
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, Product_AddProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ListProducts(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error) {
	out := new(ListProductResponse)
	err := c.cc.Invoke(ctx, Product_ListProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateProducts(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, Product_UpdateProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, Product_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetQuantityFromProductID(ctx context.Context, in *GetQuantityFromProductIDRequest, opts ...grpc.CallOption) (*GetQuantityFromProductIDResponse, error) {
	out := new(GetQuantityFromProductIDResponse)
	err := c.cc.Invoke(ctx, Product_GetQuantityFromProductID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetPriceofProductFromID(ctx context.Context, in *GetPriceofProductFromIDRequest, opts ...grpc.CallOption) (*GetPriceofProductFromIDResponse, error) {
	out := new(GetPriceofProductFromIDResponse)
	err := c.cc.Invoke(ctx, Product_GetPriceofProductFromID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductStockMinus(ctx context.Context, in *ProductStockMinusRequest, opts ...grpc.CallOption) (*ProductStockMinusReponse, error) {
	out := new(ProductStockMinusReponse)
	err := c.cc.Invoke(ctx, Product_ProductStockMinus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CheckProduct(ctx context.Context, in *CheckProductRequest, opts ...grpc.CallOption) (*CheckProductResponse, error) {
	out := new(CheckProductResponse)
	err := c.cc.Invoke(ctx, Product_CheckProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error)
	ListProducts(context.Context, *ListProductRequest) (*ListProductResponse, error)
	UpdateProducts(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	GetQuantityFromProductID(context.Context, *GetQuantityFromProductIDRequest) (*GetQuantityFromProductIDResponse, error)
	GetPriceofProductFromID(context.Context, *GetPriceofProductFromIDRequest) (*GetPriceofProductFromIDResponse, error)
	ProductStockMinus(context.Context, *ProductStockMinusRequest) (*ProductStockMinusReponse, error)
	CheckProduct(context.Context, *CheckProductRequest) (*CheckProductResponse, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductServer) ListProducts(context.Context, *ListProductRequest) (*ListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductServer) UpdateProducts(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProducts not implemented")
}
func (UnimplementedProductServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServer) GetQuantityFromProductID(context.Context, *GetQuantityFromProductIDRequest) (*GetQuantityFromProductIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuantityFromProductID not implemented")
}
func (UnimplementedProductServer) GetPriceofProductFromID(context.Context, *GetPriceofProductFromIDRequest) (*GetPriceofProductFromIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPriceofProductFromID not implemented")
}
func (UnimplementedProductServer) ProductStockMinus(context.Context, *ProductStockMinusRequest) (*ProductStockMinusReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductStockMinus not implemented")
}
func (UnimplementedProductServer) CheckProduct(context.Context, *CheckProductRequest) (*CheckProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProduct not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ListProducts(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_UpdateProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateProducts(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetQuantityFromProductID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuantityFromProductIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetQuantityFromProductID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetQuantityFromProductID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetQuantityFromProductID(ctx, req.(*GetQuantityFromProductIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetPriceofProductFromID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPriceofProductFromIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetPriceofProductFromID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetPriceofProductFromID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetPriceofProductFromID(ctx, req.(*GetPriceofProductFromIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductStockMinus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductStockMinusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductStockMinus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ProductStockMinus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductStockMinus(ctx, req.(*ProductStockMinusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CheckProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CheckProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CheckProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CheckProduct(ctx, req.(*CheckProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _Product_AddProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _Product_ListProducts_Handler,
		},
		{
			MethodName: "UpdateProducts",
			Handler:    _Product_UpdateProducts_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Product_DeleteProduct_Handler,
		},
		{
			MethodName: "GetQuantityFromProductID",
			Handler:    _Product_GetQuantityFromProductID_Handler,
		},
		{
			MethodName: "GetPriceofProductFromID",
			Handler:    _Product_GetPriceofProductFromID_Handler,
		},
		{
			MethodName: "ProductStockMinus",
			Handler:    _Product_ProductStockMinus_Handler,
		},
		{
			MethodName: "CheckProduct",
			Handler:    _Product_CheckProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/product/product.proto",
}
