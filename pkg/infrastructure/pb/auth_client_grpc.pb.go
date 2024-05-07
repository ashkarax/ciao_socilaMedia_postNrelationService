// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/infrastructure/pb/auth_client.proto

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	GetUserDetailsLiteForPostView(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseUserDetailsLite, error)
	CheckUserExist(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseBool, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GetUserDetailsLiteForPostView(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseUserDetailsLite, error) {
	out := new(ResponseUserDetailsLite)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/GetUserDetailsLiteForPostView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckUserExist(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseBool, error) {
	out := new(ResponseBool)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/CheckUserExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	GetUserDetailsLiteForPostView(context.Context, *RequestUserId) (*ResponseUserDetailsLite, error)
	CheckUserExist(context.Context, *RequestUserId) (*ResponseBool, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) GetUserDetailsLiteForPostView(context.Context, *RequestUserId) (*ResponseUserDetailsLite, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetailsLiteForPostView not implemented")
}
func (UnimplementedAuthServiceServer) CheckUserExist(context.Context, *RequestUserId) (*ResponseBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserExist not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_GetUserDetailsLiteForPostView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserDetailsLiteForPostView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/GetUserDetailsLiteForPostView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserDetailsLiteForPostView(ctx, req.(*RequestUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/CheckUserExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckUserExist(ctx, req.(*RequestUserId))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserDetailsLiteForPostView",
			Handler:    _AuthService_GetUserDetailsLiteForPostView_Handler,
		},
		{
			MethodName: "CheckUserExist",
			Handler:    _AuthService_CheckUserExist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/infrastructure/pb/auth_client.proto",
}
