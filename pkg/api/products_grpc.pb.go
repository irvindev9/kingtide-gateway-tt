// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: products.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProductsService_GetProducts_FullMethodName    = "/ProductsService/GetProducts"
	ProductsService_GetProduct_FullMethodName     = "/ProductsService/GetProduct"
	ProductsService_CreateProducts_FullMethodName = "/ProductsService/CreateProducts"
	ProductsService_UpdateProducts_FullMethodName = "/ProductsService/UpdateProducts"
	ProductsService_DeleteProducts_FullMethodName = "/ProductsService/DeleteProducts"
)

// ProductsServiceClient is the client API for ProductsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsServiceClient interface {
	GetProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
	GetProduct(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*Products, error)
	CreateProducts(ctx context.Context, in *Products, opts ...grpc.CallOption) (*Products, error)
	UpdateProducts(ctx context.Context, in *Products, opts ...grpc.CallOption) (*Products, error)
	DeleteProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
}

type productsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsServiceClient(cc grpc.ClientConnInterface) ProductsServiceClient {
	return &productsServiceClient{cc}
}

func (c *productsServiceClient) GetProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, ProductsService_GetProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) GetProduct(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*Products, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Products)
	err := c.cc.Invoke(ctx, ProductsService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) CreateProducts(ctx context.Context, in *Products, opts ...grpc.CallOption) (*Products, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Products)
	err := c.cc.Invoke(ctx, ProductsService_CreateProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) UpdateProducts(ctx context.Context, in *Products, opts ...grpc.CallOption) (*Products, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Products)
	err := c.cc.Invoke(ctx, ProductsService_UpdateProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) DeleteProducts(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, ProductsService_DeleteProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServiceServer is the server API for ProductsService service.
// All implementations must embed UnimplementedProductsServiceServer
// for forward compatibility.
type ProductsServiceServer interface {
	GetProducts(context.Context, *ProductsRequest) (*ProductsResponse, error)
	GetProduct(context.Context, *ProductsRequest) (*Products, error)
	CreateProducts(context.Context, *Products) (*Products, error)
	UpdateProducts(context.Context, *Products) (*Products, error)
	DeleteProducts(context.Context, *ProductsRequest) (*ProductsResponse, error)
	mustEmbedUnimplementedProductsServiceServer()
}

// UnimplementedProductsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductsServiceServer struct{}

func (UnimplementedProductsServiceServer) GetProducts(context.Context, *ProductsRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductsServiceServer) GetProduct(context.Context, *ProductsRequest) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductsServiceServer) CreateProducts(context.Context, *Products) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProducts not implemented")
}
func (UnimplementedProductsServiceServer) UpdateProducts(context.Context, *Products) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProducts not implemented")
}
func (UnimplementedProductsServiceServer) DeleteProducts(context.Context, *ProductsRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProducts not implemented")
}
func (UnimplementedProductsServiceServer) mustEmbedUnimplementedProductsServiceServer() {}
func (UnimplementedProductsServiceServer) testEmbeddedByValue()                         {}

// UnsafeProductsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsServiceServer will
// result in compilation errors.
type UnsafeProductsServiceServer interface {
	mustEmbedUnimplementedProductsServiceServer()
}

func RegisterProductsServiceServer(s grpc.ServiceRegistrar, srv ProductsServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductsService_ServiceDesc, srv)
}

func _ProductsService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductsService_GetProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetProducts(ctx, req.(*ProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductsService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetProduct(ctx, req.(*ProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_CreateProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Products)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).CreateProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductsService_CreateProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).CreateProducts(ctx, req.(*Products))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_UpdateProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Products)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).UpdateProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductsService_UpdateProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).UpdateProducts(ctx, req.(*Products))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_DeleteProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).DeleteProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductsService_DeleteProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).DeleteProducts(ctx, req.(*ProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductsService_ServiceDesc is the grpc.ServiceDesc for ProductsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductsService",
	HandlerType: (*ProductsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _ProductsService_GetProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductsService_GetProduct_Handler,
		},
		{
			MethodName: "CreateProducts",
			Handler:    _ProductsService_CreateProducts_Handler,
		},
		{
			MethodName: "UpdateProducts",
			Handler:    _ProductsService_UpdateProducts_Handler,
		},
		{
			MethodName: "DeleteProducts",
			Handler:    _ProductsService_DeleteProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products.proto",
}
