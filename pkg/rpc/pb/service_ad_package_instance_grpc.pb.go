// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: service_ad_package_instance.proto

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
	ADPackageInstanceService_CreateADPackageInstance_FullMethodName     = "/pb.ADPackageInstanceService/createADPackageInstance"
	ADPackageInstanceService_UpdateADPackageInstance_FullMethodName     = "/pb.ADPackageInstanceService/updateADPackageInstance"
	ADPackageInstanceService_FindADPackageInstance_FullMethodName       = "/pb.ADPackageInstanceService/findADPackageInstance"
	ADPackageInstanceService_FindAllADPackageInstances_FullMethodName   = "/pb.ADPackageInstanceService/findAllADPackageInstances"
	ADPackageInstanceService_DeleteADPackageInstance_FullMethodName     = "/pb.ADPackageInstanceService/deleteADPackageInstance"
	ADPackageInstanceService_CountIdleADPackageInstances_FullMethodName = "/pb.ADPackageInstanceService/countIdleADPackageInstances"
	ADPackageInstanceService_CountADPackageInstances_FullMethodName     = "/pb.ADPackageInstanceService/countADPackageInstances"
	ADPackageInstanceService_ListADPackageInstances_FullMethodName      = "/pb.ADPackageInstanceService/listADPackageInstances"
)

// ADPackageInstanceServiceClient is the client API for ADPackageInstanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ADPackageInstanceServiceClient interface {
	// 创建实例
	CreateADPackageInstance(ctx context.Context, in *CreateADPackageInstanceRequest, opts ...grpc.CallOption) (*CreateADPackageInstanceResponse, error)
	// 修改实例
	UpdateADPackageInstance(ctx context.Context, in *UpdateADPackageInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个实例
	FindADPackageInstance(ctx context.Context, in *FindADPackageInstanceRequest, opts ...grpc.CallOption) (*FindADPackageInstanceResponse, error)
	// 列出单个高防产品所有实例
	FindAllADPackageInstances(ctx context.Context, in *FindAllADPackageInstancesRequest, opts ...grpc.CallOption) (*FindAllADPackageInstancesResponse, error)
	// 删除实例
	DeleteADPackageInstance(ctx context.Context, in *DeleteADPackageInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算可购的实例数量
	CountIdleADPackageInstances(ctx context.Context, in *CountIdleADPackageInstancesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 计算实例数量
	CountADPackageInstances(ctx context.Context, in *CountADPackageInstancesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页实例
	ListADPackageInstances(ctx context.Context, in *ListADPackageInstancesRequest, opts ...grpc.CallOption) (*ListADPackageInstancesResponse, error)
}

type aDPackageInstanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewADPackageInstanceServiceClient(cc grpc.ClientConnInterface) ADPackageInstanceServiceClient {
	return &aDPackageInstanceServiceClient{cc}
}

func (c *aDPackageInstanceServiceClient) CreateADPackageInstance(ctx context.Context, in *CreateADPackageInstanceRequest, opts ...grpc.CallOption) (*CreateADPackageInstanceResponse, error) {
	out := new(CreateADPackageInstanceResponse)
	err := c.cc.Invoke(ctx, ADPackageInstanceService_CreateADPackageInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDPackageInstanceServiceClient) UpdateADPackageInstance(ctx context.Context, in *UpdateADPackageInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, ADPackageInstanceService_UpdateADPackageInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDPackageInstanceServiceClient) FindADPackageInstance(ctx context.Context, in *FindADPackageInstanceRequest, opts ...grpc.CallOption) (*FindADPackageInstanceResponse, error) {
	out := new(FindADPackageInstanceResponse)
	err := c.cc.Invoke(ctx, ADPackageInstanceService_FindADPackageInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDPackageInstanceServiceClient) FindAllADPackageInstances(ctx context.Context, in *FindAllADPackageInstancesRequest, opts ...grpc.CallOption) (*FindAllADPackageInstancesResponse, error) {
	out := new(FindAllADPackageInstancesResponse)
	err := c.cc.Invoke(ctx, ADPackageInstanceService_FindAllADPackageInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDPackageInstanceServiceClient) DeleteADPackageInstance(ctx context.Context, in *DeleteADPackageInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, ADPackageInstanceService_DeleteADPackageInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDPackageInstanceServiceClient) CountIdleADPackageInstances(ctx context.Context, in *CountIdleADPackageInstancesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, ADPackageInstanceService_CountIdleADPackageInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDPackageInstanceServiceClient) CountADPackageInstances(ctx context.Context, in *CountADPackageInstancesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, ADPackageInstanceService_CountADPackageInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aDPackageInstanceServiceClient) ListADPackageInstances(ctx context.Context, in *ListADPackageInstancesRequest, opts ...grpc.CallOption) (*ListADPackageInstancesResponse, error) {
	out := new(ListADPackageInstancesResponse)
	err := c.cc.Invoke(ctx, ADPackageInstanceService_ListADPackageInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ADPackageInstanceServiceServer is the server API for ADPackageInstanceService service.
// All implementations should embed UnimplementedADPackageInstanceServiceServer
// for forward compatibility
type ADPackageInstanceServiceServer interface {
	// 创建实例
	CreateADPackageInstance(context.Context, *CreateADPackageInstanceRequest) (*CreateADPackageInstanceResponse, error)
	// 修改实例
	UpdateADPackageInstance(context.Context, *UpdateADPackageInstanceRequest) (*RPCSuccess, error)
	// 查找单个实例
	FindADPackageInstance(context.Context, *FindADPackageInstanceRequest) (*FindADPackageInstanceResponse, error)
	// 列出单个高防产品所有实例
	FindAllADPackageInstances(context.Context, *FindAllADPackageInstancesRequest) (*FindAllADPackageInstancesResponse, error)
	// 删除实例
	DeleteADPackageInstance(context.Context, *DeleteADPackageInstanceRequest) (*RPCSuccess, error)
	// 计算可购的实例数量
	CountIdleADPackageInstances(context.Context, *CountIdleADPackageInstancesRequest) (*RPCCountResponse, error)
	// 计算实例数量
	CountADPackageInstances(context.Context, *CountADPackageInstancesRequest) (*RPCCountResponse, error)
	// 列出单页实例
	ListADPackageInstances(context.Context, *ListADPackageInstancesRequest) (*ListADPackageInstancesResponse, error)
}

// UnimplementedADPackageInstanceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedADPackageInstanceServiceServer struct {
}

func (UnimplementedADPackageInstanceServiceServer) CreateADPackageInstance(context.Context, *CreateADPackageInstanceRequest) (*CreateADPackageInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateADPackageInstance not implemented")
}
func (UnimplementedADPackageInstanceServiceServer) UpdateADPackageInstance(context.Context, *UpdateADPackageInstanceRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateADPackageInstance not implemented")
}
func (UnimplementedADPackageInstanceServiceServer) FindADPackageInstance(context.Context, *FindADPackageInstanceRequest) (*FindADPackageInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindADPackageInstance not implemented")
}
func (UnimplementedADPackageInstanceServiceServer) FindAllADPackageInstances(context.Context, *FindAllADPackageInstancesRequest) (*FindAllADPackageInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllADPackageInstances not implemented")
}
func (UnimplementedADPackageInstanceServiceServer) DeleteADPackageInstance(context.Context, *DeleteADPackageInstanceRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteADPackageInstance not implemented")
}
func (UnimplementedADPackageInstanceServiceServer) CountIdleADPackageInstances(context.Context, *CountIdleADPackageInstancesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountIdleADPackageInstances not implemented")
}
func (UnimplementedADPackageInstanceServiceServer) CountADPackageInstances(context.Context, *CountADPackageInstancesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountADPackageInstances not implemented")
}
func (UnimplementedADPackageInstanceServiceServer) ListADPackageInstances(context.Context, *ListADPackageInstancesRequest) (*ListADPackageInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListADPackageInstances not implemented")
}

// UnsafeADPackageInstanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ADPackageInstanceServiceServer will
// result in compilation errors.
type UnsafeADPackageInstanceServiceServer interface {
	mustEmbedUnimplementedADPackageInstanceServiceServer()
}

func RegisterADPackageInstanceServiceServer(s grpc.ServiceRegistrar, srv ADPackageInstanceServiceServer) {
	s.RegisterService(&ADPackageInstanceService_ServiceDesc, srv)
}

func _ADPackageInstanceService_CreateADPackageInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateADPackageInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADPackageInstanceServiceServer).CreateADPackageInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADPackageInstanceService_CreateADPackageInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADPackageInstanceServiceServer).CreateADPackageInstance(ctx, req.(*CreateADPackageInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADPackageInstanceService_UpdateADPackageInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateADPackageInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADPackageInstanceServiceServer).UpdateADPackageInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADPackageInstanceService_UpdateADPackageInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADPackageInstanceServiceServer).UpdateADPackageInstance(ctx, req.(*UpdateADPackageInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADPackageInstanceService_FindADPackageInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindADPackageInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADPackageInstanceServiceServer).FindADPackageInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADPackageInstanceService_FindADPackageInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADPackageInstanceServiceServer).FindADPackageInstance(ctx, req.(*FindADPackageInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADPackageInstanceService_FindAllADPackageInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllADPackageInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADPackageInstanceServiceServer).FindAllADPackageInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADPackageInstanceService_FindAllADPackageInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADPackageInstanceServiceServer).FindAllADPackageInstances(ctx, req.(*FindAllADPackageInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADPackageInstanceService_DeleteADPackageInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteADPackageInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADPackageInstanceServiceServer).DeleteADPackageInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADPackageInstanceService_DeleteADPackageInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADPackageInstanceServiceServer).DeleteADPackageInstance(ctx, req.(*DeleteADPackageInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADPackageInstanceService_CountIdleADPackageInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountIdleADPackageInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADPackageInstanceServiceServer).CountIdleADPackageInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADPackageInstanceService_CountIdleADPackageInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADPackageInstanceServiceServer).CountIdleADPackageInstances(ctx, req.(*CountIdleADPackageInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADPackageInstanceService_CountADPackageInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountADPackageInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADPackageInstanceServiceServer).CountADPackageInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADPackageInstanceService_CountADPackageInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADPackageInstanceServiceServer).CountADPackageInstances(ctx, req.(*CountADPackageInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ADPackageInstanceService_ListADPackageInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListADPackageInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADPackageInstanceServiceServer).ListADPackageInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ADPackageInstanceService_ListADPackageInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADPackageInstanceServiceServer).ListADPackageInstances(ctx, req.(*ListADPackageInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ADPackageInstanceService_ServiceDesc is the grpc.ServiceDesc for ADPackageInstanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ADPackageInstanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ADPackageInstanceService",
	HandlerType: (*ADPackageInstanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createADPackageInstance",
			Handler:    _ADPackageInstanceService_CreateADPackageInstance_Handler,
		},
		{
			MethodName: "updateADPackageInstance",
			Handler:    _ADPackageInstanceService_UpdateADPackageInstance_Handler,
		},
		{
			MethodName: "findADPackageInstance",
			Handler:    _ADPackageInstanceService_FindADPackageInstance_Handler,
		},
		{
			MethodName: "findAllADPackageInstances",
			Handler:    _ADPackageInstanceService_FindAllADPackageInstances_Handler,
		},
		{
			MethodName: "deleteADPackageInstance",
			Handler:    _ADPackageInstanceService_DeleteADPackageInstance_Handler,
		},
		{
			MethodName: "countIdleADPackageInstances",
			Handler:    _ADPackageInstanceService_CountIdleADPackageInstances_Handler,
		},
		{
			MethodName: "countADPackageInstances",
			Handler:    _ADPackageInstanceService_CountADPackageInstances_Handler,
		},
		{
			MethodName: "listADPackageInstances",
			Handler:    _ADPackageInstanceService_ListADPackageInstances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ad_package_instance.proto",
}
