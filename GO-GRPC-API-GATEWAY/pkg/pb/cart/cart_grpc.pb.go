// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: pkg/pb/cart/cart.proto

package cart

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
	Cart_AddToCart_FullMethodName = "/cart.Cart/AddToCart"
	Cart_GetCart_FullMethodName   = "/cart.Cart/GetCart"
)

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error) {
	out := new(AddToCartResponse)
	err := c.cc.Invoke(ctx, Cart_AddToCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error) {
	out := new(GetCartResponse)
	err := c.cc.Invoke(ctx, Cart_GetCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility
type CartServer interface {
	AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error)
	GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (UnimplementedCartServer) AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedCartServer) GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_AddToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).AddToCart(ctx, req.(*AddToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToCart",
			Handler:    _Cart_AddToCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _Cart_GetCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/cart/cart.proto",
}
