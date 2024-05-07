// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/infrastructure/pb/postnrelation.proto

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

// PostNrelServiceClient is the client API for PostNrelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostNrelServiceClient interface {
	AddNewPost(ctx context.Context, in *RequestAddPost, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error)
	GetAllPostByUser(ctx context.Context, in *RequestGetAllPosts, opts ...grpc.CallOption) (*ResponseUserPosts, error)
	DeletePost(ctx context.Context, in *RequestDeletePost, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error)
	EditPost(ctx context.Context, in *RequestEditPost, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error)
	Follow(ctx context.Context, in *RequestFollowUnFollow, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error)
	UnFollow(ctx context.Context, in *RequestFollowUnFollow, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error)
	// from auth svc
	GetCountsForUserProfile(ctx context.Context, in *RequestUserIdPnR, opts ...grpc.CallOption) (*ResponseGetCounts, error)
	GetFollowersIds(ctx context.Context, in *RequestUserIdPnR, opts ...grpc.CallOption) (*ResposneGetUsersIds, error)
	GetFollowingsIds(ctx context.Context, in *RequestUserIdPnR, opts ...grpc.CallOption) (*ResposneGetUsersIds, error)
}

type postNrelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostNrelServiceClient(cc grpc.ClientConnInterface) PostNrelServiceClient {
	return &postNrelServiceClient{cc}
}

func (c *postNrelServiceClient) AddNewPost(ctx context.Context, in *RequestAddPost, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error) {
	out := new(ResponseErrorMessageOnly)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/AddNewPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postNrelServiceClient) GetAllPostByUser(ctx context.Context, in *RequestGetAllPosts, opts ...grpc.CallOption) (*ResponseUserPosts, error) {
	out := new(ResponseUserPosts)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/GetAllPostByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postNrelServiceClient) DeletePost(ctx context.Context, in *RequestDeletePost, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error) {
	out := new(ResponseErrorMessageOnly)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postNrelServiceClient) EditPost(ctx context.Context, in *RequestEditPost, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error) {
	out := new(ResponseErrorMessageOnly)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/EditPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postNrelServiceClient) Follow(ctx context.Context, in *RequestFollowUnFollow, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error) {
	out := new(ResponseErrorMessageOnly)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postNrelServiceClient) UnFollow(ctx context.Context, in *RequestFollowUnFollow, opts ...grpc.CallOption) (*ResponseErrorMessageOnly, error) {
	out := new(ResponseErrorMessageOnly)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/UnFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postNrelServiceClient) GetCountsForUserProfile(ctx context.Context, in *RequestUserIdPnR, opts ...grpc.CallOption) (*ResponseGetCounts, error) {
	out := new(ResponseGetCounts)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/GetCountsForUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postNrelServiceClient) GetFollowersIds(ctx context.Context, in *RequestUserIdPnR, opts ...grpc.CallOption) (*ResposneGetUsersIds, error) {
	out := new(ResposneGetUsersIds)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/GetFollowersIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postNrelServiceClient) GetFollowingsIds(ctx context.Context, in *RequestUserIdPnR, opts ...grpc.CallOption) (*ResposneGetUsersIds, error) {
	out := new(ResposneGetUsersIds)
	err := c.cc.Invoke(ctx, "/postnrel_proto.PostNrelService/GetFollowingsIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostNrelServiceServer is the server API for PostNrelService service.
// All implementations must embed UnimplementedPostNrelServiceServer
// for forward compatibility
type PostNrelServiceServer interface {
	AddNewPost(context.Context, *RequestAddPost) (*ResponseErrorMessageOnly, error)
	GetAllPostByUser(context.Context, *RequestGetAllPosts) (*ResponseUserPosts, error)
	DeletePost(context.Context, *RequestDeletePost) (*ResponseErrorMessageOnly, error)
	EditPost(context.Context, *RequestEditPost) (*ResponseErrorMessageOnly, error)
	Follow(context.Context, *RequestFollowUnFollow) (*ResponseErrorMessageOnly, error)
	UnFollow(context.Context, *RequestFollowUnFollow) (*ResponseErrorMessageOnly, error)
	// from auth svc
	GetCountsForUserProfile(context.Context, *RequestUserIdPnR) (*ResponseGetCounts, error)
	GetFollowersIds(context.Context, *RequestUserIdPnR) (*ResposneGetUsersIds, error)
	GetFollowingsIds(context.Context, *RequestUserIdPnR) (*ResposneGetUsersIds, error)
	mustEmbedUnimplementedPostNrelServiceServer()
}

// UnimplementedPostNrelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostNrelServiceServer struct {
}

func (UnimplementedPostNrelServiceServer) AddNewPost(context.Context, *RequestAddPost) (*ResponseErrorMessageOnly, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewPost not implemented")
}
func (UnimplementedPostNrelServiceServer) GetAllPostByUser(context.Context, *RequestGetAllPosts) (*ResponseUserPosts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPostByUser not implemented")
}
func (UnimplementedPostNrelServiceServer) DeletePost(context.Context, *RequestDeletePost) (*ResponseErrorMessageOnly, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostNrelServiceServer) EditPost(context.Context, *RequestEditPost) (*ResponseErrorMessageOnly, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPost not implemented")
}
func (UnimplementedPostNrelServiceServer) Follow(context.Context, *RequestFollowUnFollow) (*ResponseErrorMessageOnly, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedPostNrelServiceServer) UnFollow(context.Context, *RequestFollowUnFollow) (*ResponseErrorMessageOnly, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollow not implemented")
}
func (UnimplementedPostNrelServiceServer) GetCountsForUserProfile(context.Context, *RequestUserIdPnR) (*ResponseGetCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountsForUserProfile not implemented")
}
func (UnimplementedPostNrelServiceServer) GetFollowersIds(context.Context, *RequestUserIdPnR) (*ResposneGetUsersIds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowersIds not implemented")
}
func (UnimplementedPostNrelServiceServer) GetFollowingsIds(context.Context, *RequestUserIdPnR) (*ResposneGetUsersIds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowingsIds not implemented")
}
func (UnimplementedPostNrelServiceServer) mustEmbedUnimplementedPostNrelServiceServer() {}

// UnsafePostNrelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostNrelServiceServer will
// result in compilation errors.
type UnsafePostNrelServiceServer interface {
	mustEmbedUnimplementedPostNrelServiceServer()
}

func RegisterPostNrelServiceServer(s grpc.ServiceRegistrar, srv PostNrelServiceServer) {
	s.RegisterService(&PostNrelService_ServiceDesc, srv)
}

func _PostNrelService_AddNewPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).AddNewPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/AddNewPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).AddNewPost(ctx, req.(*RequestAddPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostNrelService_GetAllPostByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetAllPosts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).GetAllPostByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/GetAllPostByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).GetAllPostByUser(ctx, req.(*RequestGetAllPosts))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostNrelService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeletePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).DeletePost(ctx, req.(*RequestDeletePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostNrelService_EditPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).EditPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/EditPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).EditPost(ctx, req.(*RequestEditPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostNrelService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFollowUnFollow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).Follow(ctx, req.(*RequestFollowUnFollow))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostNrelService_UnFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFollowUnFollow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).UnFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/UnFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).UnFollow(ctx, req.(*RequestFollowUnFollow))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostNrelService_GetCountsForUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserIdPnR)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).GetCountsForUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/GetCountsForUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).GetCountsForUserProfile(ctx, req.(*RequestUserIdPnR))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostNrelService_GetFollowersIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserIdPnR)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).GetFollowersIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/GetFollowersIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).GetFollowersIds(ctx, req.(*RequestUserIdPnR))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostNrelService_GetFollowingsIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserIdPnR)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostNrelServiceServer).GetFollowingsIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postnrel_proto.PostNrelService/GetFollowingsIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostNrelServiceServer).GetFollowingsIds(ctx, req.(*RequestUserIdPnR))
	}
	return interceptor(ctx, in, info, handler)
}

// PostNrelService_ServiceDesc is the grpc.ServiceDesc for PostNrelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostNrelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "postnrel_proto.PostNrelService",
	HandlerType: (*PostNrelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewPost",
			Handler:    _PostNrelService_AddNewPost_Handler,
		},
		{
			MethodName: "GetAllPostByUser",
			Handler:    _PostNrelService_GetAllPostByUser_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostNrelService_DeletePost_Handler,
		},
		{
			MethodName: "EditPost",
			Handler:    _PostNrelService_EditPost_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _PostNrelService_Follow_Handler,
		},
		{
			MethodName: "UnFollow",
			Handler:    _PostNrelService_UnFollow_Handler,
		},
		{
			MethodName: "GetCountsForUserProfile",
			Handler:    _PostNrelService_GetCountsForUserProfile_Handler,
		},
		{
			MethodName: "GetFollowersIds",
			Handler:    _PostNrelService_GetFollowersIds_Handler,
		},
		{
			MethodName: "GetFollowingsIds",
			Handler:    _PostNrelService_GetFollowingsIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/infrastructure/pb/postnrelation.proto",
}
