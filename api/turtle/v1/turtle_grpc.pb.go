// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.11.3
// source: api/turtle/v1/turtle.proto

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
	Turtle_SetTurtleBatch_FullMethodName = "/turtle.v1.turtle/SetTurtleBatch"
	Turtle_GetTurtleList_FullMethodName  = "/turtle.v1.turtle/GetTurtleList"
	Turtle_StartApp_FullMethodName       = "/turtle.v1.turtle/StartApp"
	Turtle_EndApp_FullMethodName         = "/turtle.v1.turtle/EndApp"
)

// TurtleClient is the client API for Turtle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TurtleClient interface {
	// 批量写入谜题
	SetTurtleBatch(ctx context.Context, in *SetTurtleBatchReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 分页读取谜题
	GetTurtleList(ctx context.Context, in *GetTurtleListReq, opts ...grpc.CallOption) (*GetTurtleListResp, error)
	// 开启互玩
	StartApp(ctx context.Context, in *StartAppReq, opts ...grpc.CallOption) (*StartAppResp, error)
	// 关闭互玩
	EndApp(ctx context.Context, in *EndAppReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type turtleClient struct {
	cc grpc.ClientConnInterface
}

func NewTurtleClient(cc grpc.ClientConnInterface) TurtleClient {
	return &turtleClient{cc}
}

func (c *turtleClient) SetTurtleBatch(ctx context.Context, in *SetTurtleBatchReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Turtle_SetTurtleBatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turtleClient) GetTurtleList(ctx context.Context, in *GetTurtleListReq, opts ...grpc.CallOption) (*GetTurtleListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTurtleListResp)
	err := c.cc.Invoke(ctx, Turtle_GetTurtleList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turtleClient) StartApp(ctx context.Context, in *StartAppReq, opts ...grpc.CallOption) (*StartAppResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartAppResp)
	err := c.cc.Invoke(ctx, Turtle_StartApp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turtleClient) EndApp(ctx context.Context, in *EndAppReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Turtle_EndApp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TurtleServer is the server API for Turtle service.
// All implementations must embed UnimplementedTurtleServer
// for forward compatibility.
type TurtleServer interface {
	// 批量写入谜题
	SetTurtleBatch(context.Context, *SetTurtleBatchReq) (*emptypb.Empty, error)
	// 分页读取谜题
	GetTurtleList(context.Context, *GetTurtleListReq) (*GetTurtleListResp, error)
	// 开启互玩
	StartApp(context.Context, *StartAppReq) (*StartAppResp, error)
	// 关闭互玩
	EndApp(context.Context, *EndAppReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedTurtleServer()
}

// UnimplementedTurtleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTurtleServer struct{}

func (UnimplementedTurtleServer) SetTurtleBatch(context.Context, *SetTurtleBatchReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTurtleBatch not implemented")
}
func (UnimplementedTurtleServer) GetTurtleList(context.Context, *GetTurtleListReq) (*GetTurtleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTurtleList not implemented")
}
func (UnimplementedTurtleServer) StartApp(context.Context, *StartAppReq) (*StartAppResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartApp not implemented")
}
func (UnimplementedTurtleServer) EndApp(context.Context, *EndAppReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndApp not implemented")
}
func (UnimplementedTurtleServer) mustEmbedUnimplementedTurtleServer() {}
func (UnimplementedTurtleServer) testEmbeddedByValue()                {}

// UnsafeTurtleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TurtleServer will
// result in compilation errors.
type UnsafeTurtleServer interface {
	mustEmbedUnimplementedTurtleServer()
}

func RegisterTurtleServer(s grpc.ServiceRegistrar, srv TurtleServer) {
	// If the following call pancis, it indicates UnimplementedTurtleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Turtle_ServiceDesc, srv)
}

func _Turtle_SetTurtleBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTurtleBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurtleServer).SetTurtleBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Turtle_SetTurtleBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurtleServer).SetTurtleBatch(ctx, req.(*SetTurtleBatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turtle_GetTurtleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTurtleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurtleServer).GetTurtleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Turtle_GetTurtleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurtleServer).GetTurtleList(ctx, req.(*GetTurtleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turtle_StartApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurtleServer).StartApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Turtle_StartApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurtleServer).StartApp(ctx, req.(*StartAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turtle_EndApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurtleServer).EndApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Turtle_EndApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurtleServer).EndApp(ctx, req.(*EndAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Turtle_ServiceDesc is the grpc.ServiceDesc for Turtle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Turtle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "turtle.v1.turtle",
	HandlerType: (*TurtleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTurtleBatch",
			Handler:    _Turtle_SetTurtleBatch_Handler,
		},
		{
			MethodName: "GetTurtleList",
			Handler:    _Turtle_GetTurtleList_Handler,
		},
		{
			MethodName: "StartApp",
			Handler:    _Turtle_StartApp_Handler,
		},
		{
			MethodName: "EndApp",
			Handler:    _Turtle_EndApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/turtle/v1/turtle.proto",
}
