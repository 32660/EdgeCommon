// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: service_dns_provider.proto

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
	DNSProviderService_CreateDNSProvider_FullMethodName                  = "/pb.DNSProviderService/createDNSProvider"
	DNSProviderService_UpdateDNSProvider_FullMethodName                  = "/pb.DNSProviderService/updateDNSProvider"
	DNSProviderService_CountAllEnabledDNSProviders_FullMethodName        = "/pb.DNSProviderService/countAllEnabledDNSProviders"
	DNSProviderService_ListEnabledDNSProviders_FullMethodName            = "/pb.DNSProviderService/listEnabledDNSProviders"
	DNSProviderService_FindAllEnabledDNSProviders_FullMethodName         = "/pb.DNSProviderService/findAllEnabledDNSProviders"
	DNSProviderService_DeleteDNSProvider_FullMethodName                  = "/pb.DNSProviderService/deleteDNSProvider"
	DNSProviderService_FindEnabledDNSProvider_FullMethodName             = "/pb.DNSProviderService/findEnabledDNSProvider"
	DNSProviderService_FindAllDNSProviderTypes_FullMethodName            = "/pb.DNSProviderService/findAllDNSProviderTypes"
	DNSProviderService_FindAllEnabledDNSProvidersWithType_FullMethodName = "/pb.DNSProviderService/findAllEnabledDNSProvidersWithType"
)

// DNSProviderServiceClient is the client API for DNSProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSProviderServiceClient interface {
	// 创建服务商
	CreateDNSProvider(ctx context.Context, in *CreateDNSProviderRequest, opts ...grpc.CallOption) (*CreateDNSProviderResponse, error)
	// 修改服务商
	UpdateDNSProvider(ctx context.Context, in *UpdateDNSProviderRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算服务商数量
	CountAllEnabledDNSProviders(ctx context.Context, in *CountAllEnabledDNSProvidersRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页服务商信息
	ListEnabledDNSProviders(ctx context.Context, in *ListEnabledDNSProvidersRequest, opts ...grpc.CallOption) (*ListEnabledDNSProvidersResponse, error)
	// 查找所有的DNS服务商
	FindAllEnabledDNSProviders(ctx context.Context, in *FindAllEnabledDNSProvidersRequest, opts ...grpc.CallOption) (*FindAllEnabledDNSProvidersResponse, error)
	// 删除服务商
	DeleteDNSProvider(ctx context.Context, in *DeleteDNSProviderRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个服务商
	FindEnabledDNSProvider(ctx context.Context, in *FindEnabledDNSProviderRequest, opts ...grpc.CallOption) (*FindEnabledDNSProviderResponse, error)
	// 取得所有服务商类型
	FindAllDNSProviderTypes(ctx context.Context, in *FindAllDNSProviderTypesRequest, opts ...grpc.CallOption) (*FindAllDNSProviderTypesResponse, error)
	// 取得某个类型的所有服务商
	FindAllEnabledDNSProvidersWithType(ctx context.Context, in *FindAllEnabledDNSProvidersWithTypeRequest, opts ...grpc.CallOption) (*FindAllEnabledDNSProvidersWithTypeResponse, error)
}

type dNSProviderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSProviderServiceClient(cc grpc.ClientConnInterface) DNSProviderServiceClient {
	return &dNSProviderServiceClient{cc}
}

func (c *dNSProviderServiceClient) CreateDNSProvider(ctx context.Context, in *CreateDNSProviderRequest, opts ...grpc.CallOption) (*CreateDNSProviderResponse, error) {
	out := new(CreateDNSProviderResponse)
	err := c.cc.Invoke(ctx, DNSProviderService_CreateDNSProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSProviderServiceClient) UpdateDNSProvider(ctx context.Context, in *UpdateDNSProviderRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, DNSProviderService_UpdateDNSProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSProviderServiceClient) CountAllEnabledDNSProviders(ctx context.Context, in *CountAllEnabledDNSProvidersRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, DNSProviderService_CountAllEnabledDNSProviders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSProviderServiceClient) ListEnabledDNSProviders(ctx context.Context, in *ListEnabledDNSProvidersRequest, opts ...grpc.CallOption) (*ListEnabledDNSProvidersResponse, error) {
	out := new(ListEnabledDNSProvidersResponse)
	err := c.cc.Invoke(ctx, DNSProviderService_ListEnabledDNSProviders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSProviderServiceClient) FindAllEnabledDNSProviders(ctx context.Context, in *FindAllEnabledDNSProvidersRequest, opts ...grpc.CallOption) (*FindAllEnabledDNSProvidersResponse, error) {
	out := new(FindAllEnabledDNSProvidersResponse)
	err := c.cc.Invoke(ctx, DNSProviderService_FindAllEnabledDNSProviders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSProviderServiceClient) DeleteDNSProvider(ctx context.Context, in *DeleteDNSProviderRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, DNSProviderService_DeleteDNSProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSProviderServiceClient) FindEnabledDNSProvider(ctx context.Context, in *FindEnabledDNSProviderRequest, opts ...grpc.CallOption) (*FindEnabledDNSProviderResponse, error) {
	out := new(FindEnabledDNSProviderResponse)
	err := c.cc.Invoke(ctx, DNSProviderService_FindEnabledDNSProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSProviderServiceClient) FindAllDNSProviderTypes(ctx context.Context, in *FindAllDNSProviderTypesRequest, opts ...grpc.CallOption) (*FindAllDNSProviderTypesResponse, error) {
	out := new(FindAllDNSProviderTypesResponse)
	err := c.cc.Invoke(ctx, DNSProviderService_FindAllDNSProviderTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSProviderServiceClient) FindAllEnabledDNSProvidersWithType(ctx context.Context, in *FindAllEnabledDNSProvidersWithTypeRequest, opts ...grpc.CallOption) (*FindAllEnabledDNSProvidersWithTypeResponse, error) {
	out := new(FindAllEnabledDNSProvidersWithTypeResponse)
	err := c.cc.Invoke(ctx, DNSProviderService_FindAllEnabledDNSProvidersWithType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSProviderServiceServer is the server API for DNSProviderService service.
// All implementations should embed UnimplementedDNSProviderServiceServer
// for forward compatibility
type DNSProviderServiceServer interface {
	// 创建服务商
	CreateDNSProvider(context.Context, *CreateDNSProviderRequest) (*CreateDNSProviderResponse, error)
	// 修改服务商
	UpdateDNSProvider(context.Context, *UpdateDNSProviderRequest) (*RPCSuccess, error)
	// 计算服务商数量
	CountAllEnabledDNSProviders(context.Context, *CountAllEnabledDNSProvidersRequest) (*RPCCountResponse, error)
	// 列出单页服务商信息
	ListEnabledDNSProviders(context.Context, *ListEnabledDNSProvidersRequest) (*ListEnabledDNSProvidersResponse, error)
	// 查找所有的DNS服务商
	FindAllEnabledDNSProviders(context.Context, *FindAllEnabledDNSProvidersRequest) (*FindAllEnabledDNSProvidersResponse, error)
	// 删除服务商
	DeleteDNSProvider(context.Context, *DeleteDNSProviderRequest) (*RPCSuccess, error)
	// 查找单个服务商
	FindEnabledDNSProvider(context.Context, *FindEnabledDNSProviderRequest) (*FindEnabledDNSProviderResponse, error)
	// 取得所有服务商类型
	FindAllDNSProviderTypes(context.Context, *FindAllDNSProviderTypesRequest) (*FindAllDNSProviderTypesResponse, error)
	// 取得某个类型的所有服务商
	FindAllEnabledDNSProvidersWithType(context.Context, *FindAllEnabledDNSProvidersWithTypeRequest) (*FindAllEnabledDNSProvidersWithTypeResponse, error)
}

// UnimplementedDNSProviderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDNSProviderServiceServer struct {
}

func (UnimplementedDNSProviderServiceServer) CreateDNSProvider(context.Context, *CreateDNSProviderRequest) (*CreateDNSProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDNSProvider not implemented")
}
func (UnimplementedDNSProviderServiceServer) UpdateDNSProvider(context.Context, *UpdateDNSProviderRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDNSProvider not implemented")
}
func (UnimplementedDNSProviderServiceServer) CountAllEnabledDNSProviders(context.Context, *CountAllEnabledDNSProvidersRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledDNSProviders not implemented")
}
func (UnimplementedDNSProviderServiceServer) ListEnabledDNSProviders(context.Context, *ListEnabledDNSProvidersRequest) (*ListEnabledDNSProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledDNSProviders not implemented")
}
func (UnimplementedDNSProviderServiceServer) FindAllEnabledDNSProviders(context.Context, *FindAllEnabledDNSProvidersRequest) (*FindAllEnabledDNSProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledDNSProviders not implemented")
}
func (UnimplementedDNSProviderServiceServer) DeleteDNSProvider(context.Context, *DeleteDNSProviderRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDNSProvider not implemented")
}
func (UnimplementedDNSProviderServiceServer) FindEnabledDNSProvider(context.Context, *FindEnabledDNSProviderRequest) (*FindEnabledDNSProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledDNSProvider not implemented")
}
func (UnimplementedDNSProviderServiceServer) FindAllDNSProviderTypes(context.Context, *FindAllDNSProviderTypesRequest) (*FindAllDNSProviderTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDNSProviderTypes not implemented")
}
func (UnimplementedDNSProviderServiceServer) FindAllEnabledDNSProvidersWithType(context.Context, *FindAllEnabledDNSProvidersWithTypeRequest) (*FindAllEnabledDNSProvidersWithTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledDNSProvidersWithType not implemented")
}

// UnsafeDNSProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DNSProviderServiceServer will
// result in compilation errors.
type UnsafeDNSProviderServiceServer interface {
	mustEmbedUnimplementedDNSProviderServiceServer()
}

func RegisterDNSProviderServiceServer(s grpc.ServiceRegistrar, srv DNSProviderServiceServer) {
	s.RegisterService(&DNSProviderService_ServiceDesc, srv)
}

func _DNSProviderService_CreateDNSProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDNSProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).CreateDNSProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_CreateDNSProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).CreateDNSProvider(ctx, req.(*CreateDNSProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSProviderService_UpdateDNSProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDNSProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).UpdateDNSProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_UpdateDNSProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).UpdateDNSProvider(ctx, req.(*UpdateDNSProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSProviderService_CountAllEnabledDNSProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledDNSProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).CountAllEnabledDNSProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_CountAllEnabledDNSProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).CountAllEnabledDNSProviders(ctx, req.(*CountAllEnabledDNSProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSProviderService_ListEnabledDNSProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledDNSProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).ListEnabledDNSProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_ListEnabledDNSProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).ListEnabledDNSProviders(ctx, req.(*ListEnabledDNSProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSProviderService_FindAllEnabledDNSProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledDNSProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).FindAllEnabledDNSProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_FindAllEnabledDNSProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).FindAllEnabledDNSProviders(ctx, req.(*FindAllEnabledDNSProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSProviderService_DeleteDNSProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDNSProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).DeleteDNSProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_DeleteDNSProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).DeleteDNSProvider(ctx, req.(*DeleteDNSProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSProviderService_FindEnabledDNSProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledDNSProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).FindEnabledDNSProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_FindEnabledDNSProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).FindEnabledDNSProvider(ctx, req.(*FindEnabledDNSProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSProviderService_FindAllDNSProviderTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllDNSProviderTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).FindAllDNSProviderTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_FindAllDNSProviderTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).FindAllDNSProviderTypes(ctx, req.(*FindAllDNSProviderTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSProviderService_FindAllEnabledDNSProvidersWithType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledDNSProvidersWithTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSProviderServiceServer).FindAllEnabledDNSProvidersWithType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSProviderService_FindAllEnabledDNSProvidersWithType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSProviderServiceServer).FindAllEnabledDNSProvidersWithType(ctx, req.(*FindAllEnabledDNSProvidersWithTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DNSProviderService_ServiceDesc is the grpc.ServiceDesc for DNSProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNSProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DNSProviderService",
	HandlerType: (*DNSProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createDNSProvider",
			Handler:    _DNSProviderService_CreateDNSProvider_Handler,
		},
		{
			MethodName: "updateDNSProvider",
			Handler:    _DNSProviderService_UpdateDNSProvider_Handler,
		},
		{
			MethodName: "countAllEnabledDNSProviders",
			Handler:    _DNSProviderService_CountAllEnabledDNSProviders_Handler,
		},
		{
			MethodName: "listEnabledDNSProviders",
			Handler:    _DNSProviderService_ListEnabledDNSProviders_Handler,
		},
		{
			MethodName: "findAllEnabledDNSProviders",
			Handler:    _DNSProviderService_FindAllEnabledDNSProviders_Handler,
		},
		{
			MethodName: "deleteDNSProvider",
			Handler:    _DNSProviderService_DeleteDNSProvider_Handler,
		},
		{
			MethodName: "findEnabledDNSProvider",
			Handler:    _DNSProviderService_FindEnabledDNSProvider_Handler,
		},
		{
			MethodName: "findAllDNSProviderTypes",
			Handler:    _DNSProviderService_FindAllDNSProviderTypes_Handler,
		},
		{
			MethodName: "findAllEnabledDNSProvidersWithType",
			Handler:    _DNSProviderService_FindAllEnabledDNSProvidersWithType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_dns_provider.proto",
}
