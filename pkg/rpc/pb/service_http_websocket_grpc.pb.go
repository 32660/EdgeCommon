// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: service_http_websocket.proto

package pb

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
	HTTPWebsocketService_CreateHTTPWebsocket_FullMethodName = "/pb.HTTPWebsocketService/createHTTPWebsocket"
	HTTPWebsocketService_UpdateHTTPWebsocket_FullMethodName = "/pb.HTTPWebsocketService/updateHTTPWebsocket"
)

// HTTPWebsocketServiceClient is the client API for HTTPWebsocketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPWebsocketServiceClient interface {
	// 创建Websocket配置
	CreateHTTPWebsocket(ctx context.Context, in *CreateHTTPWebsocketRequest, opts ...grpc.CallOption) (*CreateHTTPWebsocketResponse, error)
	// 修改Websocket配置
	UpdateHTTPWebsocket(ctx context.Context, in *UpdateHTTPWebsocketRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type hTTPWebsocketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPWebsocketServiceClient(cc grpc.ClientConnInterface) HTTPWebsocketServiceClient {
	return &hTTPWebsocketServiceClient{cc}
}

func (c *hTTPWebsocketServiceClient) CreateHTTPWebsocket(ctx context.Context, in *CreateHTTPWebsocketRequest, opts ...grpc.CallOption) (*CreateHTTPWebsocketResponse, error) {
	out := new(CreateHTTPWebsocketResponse)
	err := c.cc.Invoke(ctx, HTTPWebsocketService_CreateHTTPWebsocket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPWebsocketServiceClient) UpdateHTTPWebsocket(ctx context.Context, in *UpdateHTTPWebsocketRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, HTTPWebsocketService_UpdateHTTPWebsocket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPWebsocketServiceServer is the server API for HTTPWebsocketService service.
// All implementations should embed UnimplementedHTTPWebsocketServiceServer
// for forward compatibility
type HTTPWebsocketServiceServer interface {
	// 创建Websocket配置
	CreateHTTPWebsocket(context.Context, *CreateHTTPWebsocketRequest) (*CreateHTTPWebsocketResponse, error)
	// 修改Websocket配置
	UpdateHTTPWebsocket(context.Context, *UpdateHTTPWebsocketRequest) (*RPCSuccess, error)
}

// UnimplementedHTTPWebsocketServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHTTPWebsocketServiceServer struct {
}

func (UnimplementedHTTPWebsocketServiceServer) CreateHTTPWebsocket(context.Context, *CreateHTTPWebsocketRequest) (*CreateHTTPWebsocketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPWebsocket not implemented")
}
func (UnimplementedHTTPWebsocketServiceServer) UpdateHTTPWebsocket(context.Context, *UpdateHTTPWebsocketRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPWebsocket not implemented")
}

// UnsafeHTTPWebsocketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPWebsocketServiceServer will
// result in compilation errors.
type UnsafeHTTPWebsocketServiceServer interface {
	mustEmbedUnimplementedHTTPWebsocketServiceServer()
}

func RegisterHTTPWebsocketServiceServer(s grpc.ServiceRegistrar, srv HTTPWebsocketServiceServer) {
	s.RegisterService(&HTTPWebsocketService_ServiceDesc, srv)
}

func _HTTPWebsocketService_CreateHTTPWebsocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPWebsocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPWebsocketServiceServer).CreateHTTPWebsocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPWebsocketService_CreateHTTPWebsocket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPWebsocketServiceServer).CreateHTTPWebsocket(ctx, req.(*CreateHTTPWebsocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPWebsocketService_UpdateHTTPWebsocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPWebsocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPWebsocketServiceServer).UpdateHTTPWebsocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPWebsocketService_UpdateHTTPWebsocket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPWebsocketServiceServer).UpdateHTTPWebsocket(ctx, req.(*UpdateHTTPWebsocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPWebsocketService_ServiceDesc is the grpc.ServiceDesc for HTTPWebsocketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPWebsocketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPWebsocketService",
	HandlerType: (*HTTPWebsocketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPWebsocket",
			Handler:    _HTTPWebsocketService_CreateHTTPWebsocket_Handler,
		},
		{
			MethodName: "updateHTTPWebsocket",
			Handler:    _HTTPWebsocketService_UpdateHTTPWebsocket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_websocket.proto",
}
