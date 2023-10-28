// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: zerorpc.proto

package zerorpc

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

// ZerorpcClient is the client API for Zerorpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZerorpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type zerorpcClient struct {
	cc grpc.ClientConnInterface
}

func NewZerorpcClient(cc grpc.ClientConnInterface) ZerorpcClient {
	return &zerorpcClient{cc}
}

func (c *zerorpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/zerorpc.Zerorpc/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZerorpcServer is the server API for Zerorpc service.
// All implementations must embed UnimplementedZerorpcServer
// for forward compatibility
type ZerorpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedZerorpcServer()
}

// UnimplementedZerorpcServer must be embedded to have forward compatible implementations.
type UnimplementedZerorpcServer struct {
}

func (UnimplementedZerorpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedZerorpcServer) mustEmbedUnimplementedZerorpcServer() {}

// UnsafeZerorpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZerorpcServer will
// result in compilation errors.
type UnsafeZerorpcServer interface {
	mustEmbedUnimplementedZerorpcServer()
}

func RegisterZerorpcServer(s grpc.ServiceRegistrar, srv ZerorpcServer) {
	s.RegisterService(&Zerorpc_ServiceDesc, srv)
}

func _Zerorpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZerorpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zerorpc.Zerorpc/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZerorpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Zerorpc_ServiceDesc is the grpc.ServiceDesc for Zerorpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Zerorpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zerorpc.Zerorpc",
	HandlerType: (*ZerorpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Zerorpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zerorpc.proto",
}
