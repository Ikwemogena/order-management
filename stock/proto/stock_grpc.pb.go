// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc1
// source: proto/stock.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StockService_CheckStock_FullMethodName  = "/stock.StockService/CheckStock"
	StockService_UpdateStock_FullMethodName = "/stock.StockService/UpdateStock"
	StockService_CreateStock_FullMethodName = "/stock.StockService/CreateStock"
	StockService_GetStock_FullMethodName    = "/stock.StockService/GetStock"
)

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	CheckStock(ctx context.Context, in *CheckStockRequest, opts ...grpc.CallOption) (*CheckStockResponse, error)
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error)
	CreateStock(ctx context.Context, in *CreateStockRequest, opts ...grpc.CallOption) (*CreateStockResponse, error)
	GetStock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetStockResponse], error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) CheckStock(ctx context.Context, in *CheckStockRequest, opts ...grpc.CallOption) (*CheckStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckStockResponse)
	err := c.cc.Invoke(ctx, StockService_CheckStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStockResponse)
	err := c.cc.Invoke(ctx, StockService_UpdateStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) CreateStock(ctx context.Context, in *CreateStockRequest, opts ...grpc.CallOption) (*CreateStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStockResponse)
	err := c.cc.Invoke(ctx, StockService_CreateStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) GetStock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetStockResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StockService_ServiceDesc.Streams[0], StockService_GetStock_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, GetStockResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StockService_GetStockClient = grpc.ServerStreamingClient[GetStockResponse]

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility.
type StockServiceServer interface {
	CheckStock(context.Context, *CheckStockRequest) (*CheckStockResponse, error)
	UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error)
	CreateStock(context.Context, *CreateStockRequest) (*CreateStockResponse, error)
	GetStock(*emptypb.Empty, grpc.ServerStreamingServer[GetStockResponse]) error
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStockServiceServer struct{}

func (UnimplementedStockServiceServer) CheckStock(context.Context, *CheckStockRequest) (*CheckStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStock not implemented")
}
func (UnimplementedStockServiceServer) UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedStockServiceServer) CreateStock(context.Context, *CreateStockRequest) (*CreateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStock not implemented")
}
func (UnimplementedStockServiceServer) GetStock(*emptypb.Empty, grpc.ServerStreamingServer[GetStockResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}
func (UnimplementedStockServiceServer) testEmbeddedByValue()                      {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	// If the following call pancis, it indicates UnimplementedStockServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_CheckStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).CheckStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_CheckStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).CheckStock(ctx, req.(*CheckStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_UpdateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_CreateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).CreateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_CreateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).CreateStock(ctx, req.(*CreateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_GetStock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StockServiceServer).GetStock(m, &grpc.GenericServerStream[emptypb.Empty, GetStockResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StockService_GetStockServer = grpc.ServerStreamingServer[GetStockResponse]

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stock.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckStock",
			Handler:    _StockService_CheckStock_Handler,
		},
		{
			MethodName: "UpdateStock",
			Handler:    _StockService_UpdateStock_Handler,
		},
		{
			MethodName: "CreateStock",
			Handler:    _StockService_CreateStock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStock",
			Handler:       _StockService_GetStock_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/stock.proto",
}
