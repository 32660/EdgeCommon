// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: service_login.proto

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
	LoginService_FindEnabledLogin_FullMethodName = "/pb.LoginService/findEnabledLogin"
	LoginService_UpdateLogin_FullMethodName      = "/pb.LoginService/updateLogin"
)

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	// 查找认证
	FindEnabledLogin(ctx context.Context, in *FindEnabledLoginRequest, opts ...grpc.CallOption) (*FindEnabledLoginResponse, error)
	// 修改认证
	UpdateLogin(ctx context.Context, in *UpdateLoginRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) FindEnabledLogin(ctx context.Context, in *FindEnabledLoginRequest, opts ...grpc.CallOption) (*FindEnabledLoginResponse, error) {
	out := new(FindEnabledLoginResponse)
	err := c.cc.Invoke(ctx, LoginService_FindEnabledLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) UpdateLogin(ctx context.Context, in *UpdateLoginRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, LoginService_UpdateLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations should embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	// 查找认证
	FindEnabledLogin(context.Context, *FindEnabledLoginRequest) (*FindEnabledLoginResponse, error)
	// 修改认证
	UpdateLogin(context.Context, *UpdateLoginRequest) (*RPCSuccess, error)
}

// UnimplementedLoginServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) FindEnabledLogin(context.Context, *FindEnabledLoginRequest) (*FindEnabledLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledLogin not implemented")
}
func (UnimplementedLoginServiceServer) UpdateLogin(context.Context, *UpdateLoginRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLogin not implemented")
}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_FindEnabledLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).FindEnabledLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_FindEnabledLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).FindEnabledLogin(ctx, req.(*FindEnabledLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_UpdateLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).UpdateLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_UpdateLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).UpdateLogin(ctx, req.(*UpdateLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findEnabledLogin",
			Handler:    _LoginService_FindEnabledLogin_Handler,
		},
		{
			MethodName: "updateLogin",
			Handler:    _LoginService_UpdateLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_login.proto",
}
