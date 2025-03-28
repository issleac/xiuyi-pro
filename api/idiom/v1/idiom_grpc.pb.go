// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.11.3
// source: idiom/v1/idiom.proto

package v1

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
	Idiom_SetIdiomBatch_FullMethodName = "/idiom.v1.idiom/SetIdiomBatch"
	Idiom_GetIdiom_FullMethodName      = "/idiom.v1.idiom/GetIdiom"
	Idiom_GetRanking_FullMethodName    = "/idiom.v1.idiom/GetRanking"
	Idiom_UpdateRanking_FullMethodName = "/idiom.v1.idiom/UpdateRanking"
)

// IdiomClient is the client API for Idiom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdiomClient interface {
	// 批量写入成语
	SetIdiomBatch(ctx context.Context, in *SetIdiomBatchReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取成语
	GetIdiom(ctx context.Context, in *GetIdiomReq, opts ...grpc.CallOption) (*GetIdiomResp, error)
	// 获取排行榜
	GetRanking(ctx context.Context, in *GetRankingReq, opts ...grpc.CallOption) (*GetRankingResp, error)
	// 更新排行榜
	UpdateRanking(ctx context.Context, in *UpdateRankingReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type idiomClient struct {
	cc grpc.ClientConnInterface
}

func NewIdiomClient(cc grpc.ClientConnInterface) IdiomClient {
	return &idiomClient{cc}
}

func (c *idiomClient) SetIdiomBatch(ctx context.Context, in *SetIdiomBatchReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Idiom_SetIdiomBatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idiomClient) GetIdiom(ctx context.Context, in *GetIdiomReq, opts ...grpc.CallOption) (*GetIdiomResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIdiomResp)
	err := c.cc.Invoke(ctx, Idiom_GetIdiom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idiomClient) GetRanking(ctx context.Context, in *GetRankingReq, opts ...grpc.CallOption) (*GetRankingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRankingResp)
	err := c.cc.Invoke(ctx, Idiom_GetRanking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idiomClient) UpdateRanking(ctx context.Context, in *UpdateRankingReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Idiom_UpdateRanking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdiomServer is the server API for Idiom service.
// All implementations must embed UnimplementedIdiomServer
// for forward compatibility.
type IdiomServer interface {
	// 批量写入成语
	SetIdiomBatch(context.Context, *SetIdiomBatchReq) (*emptypb.Empty, error)
	// 获取成语
	GetIdiom(context.Context, *GetIdiomReq) (*GetIdiomResp, error)
	// 获取排行榜
	GetRanking(context.Context, *GetRankingReq) (*GetRankingResp, error)
	// 更新排行榜
	UpdateRanking(context.Context, *UpdateRankingReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedIdiomServer()
}

// UnimplementedIdiomServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIdiomServer struct{}

func (UnimplementedIdiomServer) SetIdiomBatch(context.Context, *SetIdiomBatchReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIdiomBatch not implemented")
}
func (UnimplementedIdiomServer) GetIdiom(context.Context, *GetIdiomReq) (*GetIdiomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdiom not implemented")
}
func (UnimplementedIdiomServer) GetRanking(context.Context, *GetRankingReq) (*GetRankingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRanking not implemented")
}
func (UnimplementedIdiomServer) UpdateRanking(context.Context, *UpdateRankingReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRanking not implemented")
}
func (UnimplementedIdiomServer) mustEmbedUnimplementedIdiomServer() {}
func (UnimplementedIdiomServer) testEmbeddedByValue()               {}

// UnsafeIdiomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdiomServer will
// result in compilation errors.
type UnsafeIdiomServer interface {
	mustEmbedUnimplementedIdiomServer()
}

func RegisterIdiomServer(s grpc.ServiceRegistrar, srv IdiomServer) {
	// If the following call pancis, it indicates UnimplementedIdiomServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Idiom_ServiceDesc, srv)
}

func _Idiom_SetIdiomBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetIdiomBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdiomServer).SetIdiomBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Idiom_SetIdiomBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdiomServer).SetIdiomBatch(ctx, req.(*SetIdiomBatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idiom_GetIdiom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdiomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdiomServer).GetIdiom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Idiom_GetIdiom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdiomServer).GetIdiom(ctx, req.(*GetIdiomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idiom_GetRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRankingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdiomServer).GetRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Idiom_GetRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdiomServer).GetRanking(ctx, req.(*GetRankingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idiom_UpdateRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRankingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdiomServer).UpdateRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Idiom_UpdateRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdiomServer).UpdateRanking(ctx, req.(*UpdateRankingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Idiom_ServiceDesc is the grpc.ServiceDesc for Idiom service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Idiom_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idiom.v1.idiom",
	HandlerType: (*IdiomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIdiomBatch",
			Handler:    _Idiom_SetIdiomBatch_Handler,
		},
		{
			MethodName: "GetIdiom",
			Handler:    _Idiom_GetIdiom_Handler,
		},
		{
			MethodName: "GetRanking",
			Handler:    _Idiom_GetRanking_Handler,
		},
		{
			MethodName: "UpdateRanking",
			Handler:    _Idiom_UpdateRanking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idiom/v1/idiom.proto",
}
