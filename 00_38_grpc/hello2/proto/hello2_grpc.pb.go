// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hello2

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

// Hello2Client is the client API for Hello2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Hello2Client interface {
	SayName(ctx context.Context, in *MsgReq, opts ...grpc.CallOption) (*MsgRsp, error)
}

type hello2Client struct {
	cc grpc.ClientConnInterface
}

func NewHello2Client(cc grpc.ClientConnInterface) Hello2Client {
	return &hello2Client{cc}
}

func (c *hello2Client) SayName(ctx context.Context, in *MsgReq, opts ...grpc.CallOption) (*MsgRsp, error) {
	out := new(MsgRsp)
	err := c.cc.Invoke(ctx, "/hello2.Hello2/SayName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Hello2Server is the server API for Hello2 service.
// All implementations must embed UnimplementedHello2Server
// for forward compatibility
type Hello2Server interface {
	SayName(context.Context, *MsgReq) (*MsgRsp, error)
	mustEmbedUnimplementedHello2Server()
}

// UnimplementedHello2Server must be embedded to have forward compatible implementations.
type UnimplementedHello2Server struct {
}

func (UnimplementedHello2Server) SayName(context.Context, *MsgReq) (*MsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayName not implemented")
}
func (UnimplementedHello2Server) mustEmbedUnimplementedHello2Server() {}

// UnsafeHello2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Hello2Server will
// result in compilation errors.
type UnsafeHello2Server interface {
	mustEmbedUnimplementedHello2Server()
}

func RegisterHello2Server(s grpc.ServiceRegistrar, srv Hello2Server) {
	s.RegisterService(&Hello2_ServiceDesc, srv)
}

func _Hello2_SayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Hello2Server).SayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello2.Hello2/SayName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Hello2Server).SayName(ctx, req.(*MsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Hello2_ServiceDesc is the grpc.ServiceDesc for Hello2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hello2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello2.Hello2",
	HandlerType: (*Hello2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayName",
			Handler:    _Hello2_SayName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello2.proto",
}
