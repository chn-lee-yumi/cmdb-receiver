// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: proto/receiver.proto

package receiver

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

// ReceiverClient is the client API for Receiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReceiverClient interface {
	// 定义Post方法，接受Data消息， 并返回Reply消息
	Post(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Reply, error)
}

type receiverClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiverClient(cc grpc.ClientConnInterface) ReceiverClient {
	return &receiverClient{cc}
}

func (c *receiverClient) Post(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Receiver/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiverServer is the server API for Receiver service.
// All implementations must embed UnimplementedReceiverServer
// for forward compatibility
type ReceiverServer interface {
	// 定义Post方法，接受Data消息， 并返回Reply消息
	Post(context.Context, *Data) (*Reply, error)
	mustEmbedUnimplementedReceiverServer()
}

// UnimplementedReceiverServer must be embedded to have forward compatible implementations.
type UnimplementedReceiverServer struct {
}

func (UnimplementedReceiverServer) Post(context.Context, *Data) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedReceiverServer) mustEmbedUnimplementedReceiverServer() {}

// UnsafeReceiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReceiverServer will
// result in compilation errors.
type UnsafeReceiverServer interface {
	mustEmbedUnimplementedReceiverServer()
}

func RegisterReceiverServer(s grpc.ServiceRegistrar, srv ReceiverServer) {
	s.RegisterService(&Receiver_ServiceDesc, srv)
}

func _Receiver_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Receiver/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverServer).Post(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

// Receiver_ServiceDesc is the grpc.ServiceDesc for Receiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Receiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Receiver",
	HandlerType: (*ReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Receiver_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/receiver.proto",
}
