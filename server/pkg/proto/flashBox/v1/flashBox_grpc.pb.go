// # Copyright 2023 LightTree <alwaysday1@qq.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/alwaysday1/flashBox.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: flashBox/v1/flashBox.proto

package v1

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

// FlashBoxClient is the client API for FlashBox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlashBoxClient interface {
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
}

type flashBoxClient struct {
	cc grpc.ClientConnInterface
}

func NewFlashBoxClient(cc grpc.ClientConnInterface) FlashBoxClient {
	return &flashBoxClient{cc}
}

func (c *flashBoxClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/v1.FlashBox/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlashBoxServer is the server API for FlashBox service.
// All implementations must embed UnimplementedFlashBoxServer
// for forward compatibility
type FlashBoxServer interface {
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	mustEmbedUnimplementedFlashBoxServer()
}

// UnimplementedFlashBoxServer must be embedded to have forward compatible implementations.
type UnimplementedFlashBoxServer struct {
}

func (UnimplementedFlashBoxServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedFlashBoxServer) mustEmbedUnimplementedFlashBoxServer() {}

// UnsafeFlashBoxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlashBoxServer will
// result in compilation errors.
type UnsafeFlashBoxServer interface {
	mustEmbedUnimplementedFlashBoxServer()
}

func RegisterFlashBoxServer(s grpc.ServiceRegistrar, srv FlashBoxServer) {
	s.RegisterService(&FlashBox_ServiceDesc, srv)
}

func _FlashBox_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashBoxServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.FlashBox/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashBoxServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlashBox_ServiceDesc is the grpc.ServiceDesc for FlashBox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlashBox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FlashBox",
	HandlerType: (*FlashBoxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _FlashBox_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flashBox/v1/flashBox.proto",
}
