// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cheese/v1/cheese.proto

package cheesev1

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

// CheeseAPIClient is the client API for CheeseAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheeseAPIClient interface {
	GetOneCheese(ctx context.Context, in *GetOneCheeseRequest, opts ...grpc.CallOption) (*GetOneCheeseResponse, error)
	GetAllCheeses(ctx context.Context, in *GetAllCheesesRequest, opts ...grpc.CallOption) (*GetAllCheesesResponse, error)
}

type cheeseAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCheeseAPIClient(cc grpc.ClientConnInterface) CheeseAPIClient {
	return &cheeseAPIClient{cc}
}

func (c *cheeseAPIClient) GetOneCheese(ctx context.Context, in *GetOneCheeseRequest, opts ...grpc.CallOption) (*GetOneCheeseResponse, error) {
	out := new(GetOneCheeseResponse)
	err := c.cc.Invoke(ctx, "/cheese.v1.CheeseAPI/GetOneCheese", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cheeseAPIClient) GetAllCheeses(ctx context.Context, in *GetAllCheesesRequest, opts ...grpc.CallOption) (*GetAllCheesesResponse, error) {
	out := new(GetAllCheesesResponse)
	err := c.cc.Invoke(ctx, "/cheese.v1.CheeseAPI/GetAllCheeses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheeseAPIServer is the server API for CheeseAPI service.
// All implementations must embed UnimplementedCheeseAPIServer
// for forward compatibility
type CheeseAPIServer interface {
	GetOneCheese(context.Context, *GetOneCheeseRequest) (*GetOneCheeseResponse, error)
	GetAllCheeses(context.Context, *GetAllCheesesRequest) (*GetAllCheesesResponse, error)
	mustEmbedUnimplementedCheeseAPIServer()
}

// UnimplementedCheeseAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCheeseAPIServer struct {
}

func (UnimplementedCheeseAPIServer) GetOneCheese(context.Context, *GetOneCheeseRequest) (*GetOneCheeseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneCheese not implemented")
}
func (UnimplementedCheeseAPIServer) GetAllCheeses(context.Context, *GetAllCheesesRequest) (*GetAllCheesesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCheeses not implemented")
}
func (UnimplementedCheeseAPIServer) mustEmbedUnimplementedCheeseAPIServer() {}

// UnsafeCheeseAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheeseAPIServer will
// result in compilation errors.
type UnsafeCheeseAPIServer interface {
	mustEmbedUnimplementedCheeseAPIServer()
}

func RegisterCheeseAPIServer(s grpc.ServiceRegistrar, srv CheeseAPIServer) {
	s.RegisterService(&CheeseAPI_ServiceDesc, srv)
}

func _CheeseAPI_GetOneCheese_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneCheeseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheeseAPIServer).GetOneCheese(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheese.v1.CheeseAPI/GetOneCheese",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheeseAPIServer).GetOneCheese(ctx, req.(*GetOneCheeseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheeseAPI_GetAllCheeses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCheesesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheeseAPIServer).GetAllCheeses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheese.v1.CheeseAPI/GetAllCheeses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheeseAPIServer).GetAllCheeses(ctx, req.(*GetAllCheesesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheeseAPI_ServiceDesc is the grpc.ServiceDesc for CheeseAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheeseAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cheese.v1.CheeseAPI",
	HandlerType: (*CheeseAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOneCheese",
			Handler:    _CheeseAPI_GetOneCheese_Handler,
		},
		{
			MethodName: "GetAllCheeses",
			Handler:    _CheeseAPI_GetAllCheeses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheese/v1/cheese.proto",
}
