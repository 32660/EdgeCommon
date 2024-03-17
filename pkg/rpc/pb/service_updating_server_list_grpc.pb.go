// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: service_updating_server_list.proto

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
	UpdatingServerListService_FindUpdatingServerLists_FullMethodName = "/pb.UpdatingServerListService/findUpdatingServerLists"
)

// UpdatingServerListServiceClient is the client API for UpdatingServerListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdatingServerListServiceClient interface {
	// 查找要更新的服务配置
	FindUpdatingServerLists(ctx context.Context, in *FindUpdatingServerListsRequest, opts ...grpc.CallOption) (*FindUpdatingServerListsResponse, error)
}

type updatingServerListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdatingServerListServiceClient(cc grpc.ClientConnInterface) UpdatingServerListServiceClient {
	return &updatingServerListServiceClient{cc}
}

func (c *updatingServerListServiceClient) FindUpdatingServerLists(ctx context.Context, in *FindUpdatingServerListsRequest, opts ...grpc.CallOption) (*FindUpdatingServerListsResponse, error) {
	out := new(FindUpdatingServerListsResponse)
	err := c.cc.Invoke(ctx, UpdatingServerListService_FindUpdatingServerLists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdatingServerListServiceServer is the server API for UpdatingServerListService service.
// All implementations should embed UnimplementedUpdatingServerListServiceServer
// for forward compatibility
type UpdatingServerListServiceServer interface {
	// 查找要更新的服务配置
	FindUpdatingServerLists(context.Context, *FindUpdatingServerListsRequest) (*FindUpdatingServerListsResponse, error)
}

// UnimplementedUpdatingServerListServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUpdatingServerListServiceServer struct {
}

func (UnimplementedUpdatingServerListServiceServer) FindUpdatingServerLists(context.Context, *FindUpdatingServerListsRequest) (*FindUpdatingServerListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUpdatingServerLists not implemented")
}

// UnsafeUpdatingServerListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdatingServerListServiceServer will
// result in compilation errors.
type UnsafeUpdatingServerListServiceServer interface {
	mustEmbedUnimplementedUpdatingServerListServiceServer()
}

func RegisterUpdatingServerListServiceServer(s grpc.ServiceRegistrar, srv UpdatingServerListServiceServer) {
	s.RegisterService(&UpdatingServerListService_ServiceDesc, srv)
}

func _UpdatingServerListService_FindUpdatingServerLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUpdatingServerListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdatingServerListServiceServer).FindUpdatingServerLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UpdatingServerListService_FindUpdatingServerLists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdatingServerListServiceServer).FindUpdatingServerLists(ctx, req.(*FindUpdatingServerListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdatingServerListService_ServiceDesc is the grpc.ServiceDesc for UpdatingServerListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdatingServerListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UpdatingServerListService",
	HandlerType: (*UpdatingServerListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findUpdatingServerLists",
			Handler:    _UpdatingServerListService_FindUpdatingServerLists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_updating_server_list.proto",
}
