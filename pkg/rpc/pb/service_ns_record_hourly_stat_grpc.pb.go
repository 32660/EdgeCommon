// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: service_ns_record_hourly_stat.proto

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
	NSRecordHourlyStatService_UploadNSRecordHourlyStats_FullMethodName           = "/pb.NSRecordHourlyStatService/uploadNSRecordHourlyStats"
	NSRecordHourlyStatService_FindNSRecordHourlyStat_FullMethodName              = "/pb.NSRecordHourlyStatService/findNSRecordHourlyStat"
	NSRecordHourlyStatService_FindLatestNSRecordsHourlyStats_FullMethodName      = "/pb.NSRecordHourlyStatService/findLatestNSRecordsHourlyStats"
	NSRecordHourlyStatService_FindNSRecordHourlyStatWithRecordIds_FullMethodName = "/pb.NSRecordHourlyStatService/findNSRecordHourlyStatWithRecordIds"
)

// NSRecordHourlyStatServiceClient is the client API for NSRecordHourlyStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NSRecordHourlyStatServiceClient interface {
	// 上传统计
	UploadNSRecordHourlyStats(ctx context.Context, in *UploadNSRecordHourlyStatsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 获取单个记录单个小时的统计
	FindNSRecordHourlyStat(ctx context.Context, in *FindNSRecordHourlyStatRequest, opts ...grpc.CallOption) (*FindNSRecordHourlyStatResponse, error)
	// 获取单个记录24小时内的统计
	FindLatestNSRecordsHourlyStats(ctx context.Context, in *FindLatestNSRecordsHourlyStatsRequest, opts ...grpc.CallOption) (*FindLatestNSRecordsHourlyStatsResponse, error)
	// 批量获取一组记录的统计
	FindNSRecordHourlyStatWithRecordIds(ctx context.Context, in *FindNSRecordHourlyStatWithRecordIdsRequest, opts ...grpc.CallOption) (*FindNSRecordHourlyStatWithRecordIdsResponse, error)
}

type nSRecordHourlyStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNSRecordHourlyStatServiceClient(cc grpc.ClientConnInterface) NSRecordHourlyStatServiceClient {
	return &nSRecordHourlyStatServiceClient{cc}
}

func (c *nSRecordHourlyStatServiceClient) UploadNSRecordHourlyStats(ctx context.Context, in *UploadNSRecordHourlyStatsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSRecordHourlyStatService_UploadNSRecordHourlyStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSRecordHourlyStatServiceClient) FindNSRecordHourlyStat(ctx context.Context, in *FindNSRecordHourlyStatRequest, opts ...grpc.CallOption) (*FindNSRecordHourlyStatResponse, error) {
	out := new(FindNSRecordHourlyStatResponse)
	err := c.cc.Invoke(ctx, NSRecordHourlyStatService_FindNSRecordHourlyStat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSRecordHourlyStatServiceClient) FindLatestNSRecordsHourlyStats(ctx context.Context, in *FindLatestNSRecordsHourlyStatsRequest, opts ...grpc.CallOption) (*FindLatestNSRecordsHourlyStatsResponse, error) {
	out := new(FindLatestNSRecordsHourlyStatsResponse)
	err := c.cc.Invoke(ctx, NSRecordHourlyStatService_FindLatestNSRecordsHourlyStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSRecordHourlyStatServiceClient) FindNSRecordHourlyStatWithRecordIds(ctx context.Context, in *FindNSRecordHourlyStatWithRecordIdsRequest, opts ...grpc.CallOption) (*FindNSRecordHourlyStatWithRecordIdsResponse, error) {
	out := new(FindNSRecordHourlyStatWithRecordIdsResponse)
	err := c.cc.Invoke(ctx, NSRecordHourlyStatService_FindNSRecordHourlyStatWithRecordIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSRecordHourlyStatServiceServer is the server API for NSRecordHourlyStatService service.
// All implementations should embed UnimplementedNSRecordHourlyStatServiceServer
// for forward compatibility
type NSRecordHourlyStatServiceServer interface {
	// 上传统计
	UploadNSRecordHourlyStats(context.Context, *UploadNSRecordHourlyStatsRequest) (*RPCSuccess, error)
	// 获取单个记录单个小时的统计
	FindNSRecordHourlyStat(context.Context, *FindNSRecordHourlyStatRequest) (*FindNSRecordHourlyStatResponse, error)
	// 获取单个记录24小时内的统计
	FindLatestNSRecordsHourlyStats(context.Context, *FindLatestNSRecordsHourlyStatsRequest) (*FindLatestNSRecordsHourlyStatsResponse, error)
	// 批量获取一组记录的统计
	FindNSRecordHourlyStatWithRecordIds(context.Context, *FindNSRecordHourlyStatWithRecordIdsRequest) (*FindNSRecordHourlyStatWithRecordIdsResponse, error)
}

// UnimplementedNSRecordHourlyStatServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNSRecordHourlyStatServiceServer struct {
}

func (UnimplementedNSRecordHourlyStatServiceServer) UploadNSRecordHourlyStats(context.Context, *UploadNSRecordHourlyStatsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadNSRecordHourlyStats not implemented")
}
func (UnimplementedNSRecordHourlyStatServiceServer) FindNSRecordHourlyStat(context.Context, *FindNSRecordHourlyStatRequest) (*FindNSRecordHourlyStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSRecordHourlyStat not implemented")
}
func (UnimplementedNSRecordHourlyStatServiceServer) FindLatestNSRecordsHourlyStats(context.Context, *FindLatestNSRecordsHourlyStatsRequest) (*FindLatestNSRecordsHourlyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLatestNSRecordsHourlyStats not implemented")
}
func (UnimplementedNSRecordHourlyStatServiceServer) FindNSRecordHourlyStatWithRecordIds(context.Context, *FindNSRecordHourlyStatWithRecordIdsRequest) (*FindNSRecordHourlyStatWithRecordIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSRecordHourlyStatWithRecordIds not implemented")
}

// UnsafeNSRecordHourlyStatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NSRecordHourlyStatServiceServer will
// result in compilation errors.
type UnsafeNSRecordHourlyStatServiceServer interface {
	mustEmbedUnimplementedNSRecordHourlyStatServiceServer()
}

func RegisterNSRecordHourlyStatServiceServer(s grpc.ServiceRegistrar, srv NSRecordHourlyStatServiceServer) {
	s.RegisterService(&NSRecordHourlyStatService_ServiceDesc, srv)
}

func _NSRecordHourlyStatService_UploadNSRecordHourlyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadNSRecordHourlyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSRecordHourlyStatServiceServer).UploadNSRecordHourlyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSRecordHourlyStatService_UploadNSRecordHourlyStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSRecordHourlyStatServiceServer).UploadNSRecordHourlyStats(ctx, req.(*UploadNSRecordHourlyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSRecordHourlyStatService_FindNSRecordHourlyStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSRecordHourlyStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSRecordHourlyStatServiceServer).FindNSRecordHourlyStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSRecordHourlyStatService_FindNSRecordHourlyStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSRecordHourlyStatServiceServer).FindNSRecordHourlyStat(ctx, req.(*FindNSRecordHourlyStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSRecordHourlyStatService_FindLatestNSRecordsHourlyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLatestNSRecordsHourlyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSRecordHourlyStatServiceServer).FindLatestNSRecordsHourlyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSRecordHourlyStatService_FindLatestNSRecordsHourlyStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSRecordHourlyStatServiceServer).FindLatestNSRecordsHourlyStats(ctx, req.(*FindLatestNSRecordsHourlyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSRecordHourlyStatService_FindNSRecordHourlyStatWithRecordIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSRecordHourlyStatWithRecordIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSRecordHourlyStatServiceServer).FindNSRecordHourlyStatWithRecordIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSRecordHourlyStatService_FindNSRecordHourlyStatWithRecordIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSRecordHourlyStatServiceServer).FindNSRecordHourlyStatWithRecordIds(ctx, req.(*FindNSRecordHourlyStatWithRecordIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NSRecordHourlyStatService_ServiceDesc is the grpc.ServiceDesc for NSRecordHourlyStatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NSRecordHourlyStatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NSRecordHourlyStatService",
	HandlerType: (*NSRecordHourlyStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "uploadNSRecordHourlyStats",
			Handler:    _NSRecordHourlyStatService_UploadNSRecordHourlyStats_Handler,
		},
		{
			MethodName: "findNSRecordHourlyStat",
			Handler:    _NSRecordHourlyStatService_FindNSRecordHourlyStat_Handler,
		},
		{
			MethodName: "findLatestNSRecordsHourlyStats",
			Handler:    _NSRecordHourlyStatService_FindLatestNSRecordsHourlyStats_Handler,
		},
		{
			MethodName: "findNSRecordHourlyStatWithRecordIds",
			Handler:    _NSRecordHourlyStatService_FindNSRecordHourlyStatWithRecordIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ns_record_hourly_stat.proto",
}
