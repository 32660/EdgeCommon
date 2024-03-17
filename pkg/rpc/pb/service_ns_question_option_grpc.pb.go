// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: service_ns_question_option.proto

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
	NSQuestionOptionService_CreateNSQuestionOption_FullMethodName = "/pb.NSQuestionOptionService/createNSQuestionOption"
	NSQuestionOptionService_FindNSQuestionOption_FullMethodName   = "/pb.NSQuestionOptionService/findNSQuestionOption"
	NSQuestionOptionService_DeleteNSQuestionOption_FullMethodName = "/pb.NSQuestionOptionService/deleteNSQuestionOption"
)

// NSQuestionOptionServiceClient is the client API for NSQuestionOptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NSQuestionOptionServiceClient interface {
	// 创建选项
	CreateNSQuestionOption(ctx context.Context, in *CreateNSQuestionOptionRequest, opts ...grpc.CallOption) (*CreateNSQuestionOptionResponse, error)
	// 读取选项
	FindNSQuestionOption(ctx context.Context, in *FindNSQuestionOptionRequest, opts ...grpc.CallOption) (*FindNSQuestionOptionResponse, error)
	// 删除选项
	DeleteNSQuestionOption(ctx context.Context, in *DeleteNSQuestionOptionRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type nSQuestionOptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNSQuestionOptionServiceClient(cc grpc.ClientConnInterface) NSQuestionOptionServiceClient {
	return &nSQuestionOptionServiceClient{cc}
}

func (c *nSQuestionOptionServiceClient) CreateNSQuestionOption(ctx context.Context, in *CreateNSQuestionOptionRequest, opts ...grpc.CallOption) (*CreateNSQuestionOptionResponse, error) {
	out := new(CreateNSQuestionOptionResponse)
	err := c.cc.Invoke(ctx, NSQuestionOptionService_CreateNSQuestionOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSQuestionOptionServiceClient) FindNSQuestionOption(ctx context.Context, in *FindNSQuestionOptionRequest, opts ...grpc.CallOption) (*FindNSQuestionOptionResponse, error) {
	out := new(FindNSQuestionOptionResponse)
	err := c.cc.Invoke(ctx, NSQuestionOptionService_FindNSQuestionOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSQuestionOptionServiceClient) DeleteNSQuestionOption(ctx context.Context, in *DeleteNSQuestionOptionRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSQuestionOptionService_DeleteNSQuestionOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSQuestionOptionServiceServer is the server API for NSQuestionOptionService service.
// All implementations should embed UnimplementedNSQuestionOptionServiceServer
// for forward compatibility
type NSQuestionOptionServiceServer interface {
	// 创建选项
	CreateNSQuestionOption(context.Context, *CreateNSQuestionOptionRequest) (*CreateNSQuestionOptionResponse, error)
	// 读取选项
	FindNSQuestionOption(context.Context, *FindNSQuestionOptionRequest) (*FindNSQuestionOptionResponse, error)
	// 删除选项
	DeleteNSQuestionOption(context.Context, *DeleteNSQuestionOptionRequest) (*RPCSuccess, error)
}

// UnimplementedNSQuestionOptionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNSQuestionOptionServiceServer struct {
}

func (UnimplementedNSQuestionOptionServiceServer) CreateNSQuestionOption(context.Context, *CreateNSQuestionOptionRequest) (*CreateNSQuestionOptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNSQuestionOption not implemented")
}
func (UnimplementedNSQuestionOptionServiceServer) FindNSQuestionOption(context.Context, *FindNSQuestionOptionRequest) (*FindNSQuestionOptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSQuestionOption not implemented")
}
func (UnimplementedNSQuestionOptionServiceServer) DeleteNSQuestionOption(context.Context, *DeleteNSQuestionOptionRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNSQuestionOption not implemented")
}

// UnsafeNSQuestionOptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NSQuestionOptionServiceServer will
// result in compilation errors.
type UnsafeNSQuestionOptionServiceServer interface {
	mustEmbedUnimplementedNSQuestionOptionServiceServer()
}

func RegisterNSQuestionOptionServiceServer(s grpc.ServiceRegistrar, srv NSQuestionOptionServiceServer) {
	s.RegisterService(&NSQuestionOptionService_ServiceDesc, srv)
}

func _NSQuestionOptionService_CreateNSQuestionOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNSQuestionOptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSQuestionOptionServiceServer).CreateNSQuestionOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSQuestionOptionService_CreateNSQuestionOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSQuestionOptionServiceServer).CreateNSQuestionOption(ctx, req.(*CreateNSQuestionOptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSQuestionOptionService_FindNSQuestionOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSQuestionOptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSQuestionOptionServiceServer).FindNSQuestionOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSQuestionOptionService_FindNSQuestionOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSQuestionOptionServiceServer).FindNSQuestionOption(ctx, req.(*FindNSQuestionOptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSQuestionOptionService_DeleteNSQuestionOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNSQuestionOptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSQuestionOptionServiceServer).DeleteNSQuestionOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSQuestionOptionService_DeleteNSQuestionOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSQuestionOptionServiceServer).DeleteNSQuestionOption(ctx, req.(*DeleteNSQuestionOptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NSQuestionOptionService_ServiceDesc is the grpc.ServiceDesc for NSQuestionOptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NSQuestionOptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NSQuestionOptionService",
	HandlerType: (*NSQuestionOptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNSQuestionOption",
			Handler:    _NSQuestionOptionService_CreateNSQuestionOption_Handler,
		},
		{
			MethodName: "findNSQuestionOption",
			Handler:    _NSQuestionOptionService_FindNSQuestionOption_Handler,
		},
		{
			MethodName: "deleteNSQuestionOption",
			Handler:    _NSQuestionOptionService_DeleteNSQuestionOption_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ns_question_option.proto",
}
