// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: forum.proto

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

// ForumServiceClient is the client API for ForumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ForumServiceClient interface {
	CreateForum(ctx context.Context, in *Forum, opts ...grpc.CallOption) (*ForumIdResponse, error)
	CreateThread(ctx context.Context, in *Thread, opts ...grpc.CallOption) (*ForumIdResponse, error)
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*ForumIdResponse, error)
	GetForum(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*Forum, error)
	GetForums(ctx context.Context, in *GetForumsRequest, opts ...grpc.CallOption) (ForumService_GetForumsClient, error)
	GetThread(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*Thread, error)
	GetThreads(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (ForumService_GetThreadsClient, error)
	GetPost(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*Post, error)
	GetPosts(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (ForumService_GetPostsClient, error)
	UpdateForum(ctx context.Context, in *Forum, opts ...grpc.CallOption) (*Forum, error)
	UpdateThread(ctx context.Context, in *Thread, opts ...grpc.CallOption) (*Thread, error)
	UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	DeleteForum(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*ForumIdResponse, error)
	DeleteThread(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*ForumIdResponse, error)
	DeletePost(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*ForumIdResponse, error)
	SearchForum(ctx context.Context, in *ForumSearchRequest, opts ...grpc.CallOption) (ForumService_SearchForumClient, error)
}

type forumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewForumServiceClient(cc grpc.ClientConnInterface) ForumServiceClient {
	return &forumServiceClient{cc}
}

func (c *forumServiceClient) CreateForum(ctx context.Context, in *Forum, opts ...grpc.CallOption) (*ForumIdResponse, error) {
	out := new(ForumIdResponse)
	err := c.cc.Invoke(ctx, "/pb.ForumService/CreateForum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) CreateThread(ctx context.Context, in *Thread, opts ...grpc.CallOption) (*ForumIdResponse, error) {
	out := new(ForumIdResponse)
	err := c.cc.Invoke(ctx, "/pb.ForumService/CreateThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*ForumIdResponse, error) {
	out := new(ForumIdResponse)
	err := c.cc.Invoke(ctx, "/pb.ForumService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetForum(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*Forum, error) {
	out := new(Forum)
	err := c.cc.Invoke(ctx, "/pb.ForumService/GetForum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetForums(ctx context.Context, in *GetForumsRequest, opts ...grpc.CallOption) (ForumService_GetForumsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ForumService_ServiceDesc.Streams[0], "/pb.ForumService/GetForums", opts...)
	if err != nil {
		return nil, err
	}
	x := &forumServiceGetForumsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ForumService_GetForumsClient interface {
	Recv() (*Forum, error)
	grpc.ClientStream
}

type forumServiceGetForumsClient struct {
	grpc.ClientStream
}

func (x *forumServiceGetForumsClient) Recv() (*Forum, error) {
	m := new(Forum)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *forumServiceClient) GetThread(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*Thread, error) {
	out := new(Thread)
	err := c.cc.Invoke(ctx, "/pb.ForumService/GetThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetThreads(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (ForumService_GetThreadsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ForumService_ServiceDesc.Streams[1], "/pb.ForumService/GetThreads", opts...)
	if err != nil {
		return nil, err
	}
	x := &forumServiceGetThreadsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ForumService_GetThreadsClient interface {
	Recv() (*Thread, error)
	grpc.ClientStream
}

type forumServiceGetThreadsClient struct {
	grpc.ClientStream
}

func (x *forumServiceGetThreadsClient) Recv() (*Thread, error) {
	m := new(Thread)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *forumServiceClient) GetPost(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/pb.ForumService/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) GetPosts(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (ForumService_GetPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ForumService_ServiceDesc.Streams[2], "/pb.ForumService/GetPosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &forumServiceGetPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ForumService_GetPostsClient interface {
	Recv() (*Post, error)
	grpc.ClientStream
}

type forumServiceGetPostsClient struct {
	grpc.ClientStream
}

func (x *forumServiceGetPostsClient) Recv() (*Post, error) {
	m := new(Post)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *forumServiceClient) UpdateForum(ctx context.Context, in *Forum, opts ...grpc.CallOption) (*Forum, error) {
	out := new(Forum)
	err := c.cc.Invoke(ctx, "/pb.ForumService/UpdateForum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) UpdateThread(ctx context.Context, in *Thread, opts ...grpc.CallOption) (*Thread, error) {
	out := new(Thread)
	err := c.cc.Invoke(ctx, "/pb.ForumService/UpdateThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/pb.ForumService/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) DeleteForum(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*ForumIdResponse, error) {
	out := new(ForumIdResponse)
	err := c.cc.Invoke(ctx, "/pb.ForumService/DeleteForum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) DeleteThread(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*ForumIdResponse, error) {
	out := new(ForumIdResponse)
	err := c.cc.Invoke(ctx, "/pb.ForumService/DeleteThread", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) DeletePost(ctx context.Context, in *ForumIdRequest, opts ...grpc.CallOption) (*ForumIdResponse, error) {
	out := new(ForumIdResponse)
	err := c.cc.Invoke(ctx, "/pb.ForumService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumServiceClient) SearchForum(ctx context.Context, in *ForumSearchRequest, opts ...grpc.CallOption) (ForumService_SearchForumClient, error) {
	stream, err := c.cc.NewStream(ctx, &ForumService_ServiceDesc.Streams[3], "/pb.ForumService/SearchForum", opts...)
	if err != nil {
		return nil, err
	}
	x := &forumServiceSearchForumClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ForumService_SearchForumClient interface {
	Recv() (*Thread, error)
	grpc.ClientStream
}

type forumServiceSearchForumClient struct {
	grpc.ClientStream
}

func (x *forumServiceSearchForumClient) Recv() (*Thread, error) {
	m := new(Thread)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ForumServiceServer is the server API for ForumService service.
// All implementations must embed UnimplementedForumServiceServer
// for forward compatibility
type ForumServiceServer interface {
	CreateForum(context.Context, *Forum) (*ForumIdResponse, error)
	CreateThread(context.Context, *Thread) (*ForumIdResponse, error)
	CreatePost(context.Context, *Post) (*ForumIdResponse, error)
	GetForum(context.Context, *ForumIdRequest) (*Forum, error)
	GetForums(*GetForumsRequest, ForumService_GetForumsServer) error
	GetThread(context.Context, *ForumIdRequest) (*Thread, error)
	GetThreads(*ForumIdRequest, ForumService_GetThreadsServer) error
	GetPost(context.Context, *ForumIdRequest) (*Post, error)
	GetPosts(*ForumIdRequest, ForumService_GetPostsServer) error
	UpdateForum(context.Context, *Forum) (*Forum, error)
	UpdateThread(context.Context, *Thread) (*Thread, error)
	UpdatePost(context.Context, *Post) (*Post, error)
	DeleteForum(context.Context, *ForumIdRequest) (*ForumIdResponse, error)
	DeleteThread(context.Context, *ForumIdRequest) (*ForumIdResponse, error)
	DeletePost(context.Context, *ForumIdRequest) (*ForumIdResponse, error)
	SearchForum(*ForumSearchRequest, ForumService_SearchForumServer) error
	mustEmbedUnimplementedForumServiceServer()
}

// UnimplementedForumServiceServer must be embedded to have forward compatible implementations.
type UnimplementedForumServiceServer struct {
}

func (UnimplementedForumServiceServer) CreateForum(context.Context, *Forum) (*ForumIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateForum not implemented")
}
func (UnimplementedForumServiceServer) CreateThread(context.Context, *Thread) (*ForumIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateThread not implemented")
}
func (UnimplementedForumServiceServer) CreatePost(context.Context, *Post) (*ForumIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedForumServiceServer) GetForum(context.Context, *ForumIdRequest) (*Forum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForum not implemented")
}
func (UnimplementedForumServiceServer) GetForums(*GetForumsRequest, ForumService_GetForumsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetForums not implemented")
}
func (UnimplementedForumServiceServer) GetThread(context.Context, *ForumIdRequest) (*Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThread not implemented")
}
func (UnimplementedForumServiceServer) GetThreads(*ForumIdRequest, ForumService_GetThreadsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetThreads not implemented")
}
func (UnimplementedForumServiceServer) GetPost(context.Context, *ForumIdRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedForumServiceServer) GetPosts(*ForumIdRequest, ForumService_GetPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedForumServiceServer) UpdateForum(context.Context, *Forum) (*Forum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateForum not implemented")
}
func (UnimplementedForumServiceServer) UpdateThread(context.Context, *Thread) (*Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateThread not implemented")
}
func (UnimplementedForumServiceServer) UpdatePost(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedForumServiceServer) DeleteForum(context.Context, *ForumIdRequest) (*ForumIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteForum not implemented")
}
func (UnimplementedForumServiceServer) DeleteThread(context.Context, *ForumIdRequest) (*ForumIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteThread not implemented")
}
func (UnimplementedForumServiceServer) DeletePost(context.Context, *ForumIdRequest) (*ForumIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedForumServiceServer) SearchForum(*ForumSearchRequest, ForumService_SearchForumServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchForum not implemented")
}
func (UnimplementedForumServiceServer) mustEmbedUnimplementedForumServiceServer() {}

// UnsafeForumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForumServiceServer will
// result in compilation errors.
type UnsafeForumServiceServer interface {
	mustEmbedUnimplementedForumServiceServer()
}

func RegisterForumServiceServer(s grpc.ServiceRegistrar, srv ForumServiceServer) {
	s.RegisterService(&ForumService_ServiceDesc, srv)
}

func _ForumService_CreateForum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Forum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateForum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/CreateForum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateForum(ctx, req.(*Forum))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_CreateThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Thread)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreateThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/CreateThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreateThread(ctx, req.(*Thread))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetForum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForumIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetForum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/GetForum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetForum(ctx, req.(*ForumIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetForums_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetForumsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ForumServiceServer).GetForums(m, &forumServiceGetForumsServer{stream})
}

type ForumService_GetForumsServer interface {
	Send(*Forum) error
	grpc.ServerStream
}

type forumServiceGetForumsServer struct {
	grpc.ServerStream
}

func (x *forumServiceGetForumsServer) Send(m *Forum) error {
	return x.ServerStream.SendMsg(m)
}

func _ForumService_GetThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForumIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/GetThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetThread(ctx, req.(*ForumIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetThreads_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ForumIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ForumServiceServer).GetThreads(m, &forumServiceGetThreadsServer{stream})
}

type ForumService_GetThreadsServer interface {
	Send(*Thread) error
	grpc.ServerStream
}

type forumServiceGetThreadsServer struct {
	grpc.ServerStream
}

func (x *forumServiceGetThreadsServer) Send(m *Thread) error {
	return x.ServerStream.SendMsg(m)
}

func _ForumService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForumIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).GetPost(ctx, req.(*ForumIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_GetPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ForumIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ForumServiceServer).GetPosts(m, &forumServiceGetPostsServer{stream})
}

type ForumService_GetPostsServer interface {
	Send(*Post) error
	grpc.ServerStream
}

type forumServiceGetPostsServer struct {
	grpc.ServerStream
}

func (x *forumServiceGetPostsServer) Send(m *Post) error {
	return x.ServerStream.SendMsg(m)
}

func _ForumService_UpdateForum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Forum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).UpdateForum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/UpdateForum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).UpdateForum(ctx, req.(*Forum))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_UpdateThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Thread)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).UpdateThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/UpdateThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).UpdateThread(ctx, req.(*Thread))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).UpdatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_DeleteForum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForumIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).DeleteForum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/DeleteForum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).DeleteForum(ctx, req.(*ForumIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_DeleteThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForumIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).DeleteThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/DeleteThread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).DeleteThread(ctx, req.(*ForumIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForumIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ForumService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServiceServer).DeletePost(ctx, req.(*ForumIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForumService_SearchForum_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ForumSearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ForumServiceServer).SearchForum(m, &forumServiceSearchForumServer{stream})
}

type ForumService_SearchForumServer interface {
	Send(*Thread) error
	grpc.ServerStream
}

type forumServiceSearchForumServer struct {
	grpc.ServerStream
}

func (x *forumServiceSearchForumServer) Send(m *Thread) error {
	return x.ServerStream.SendMsg(m)
}

// ForumService_ServiceDesc is the grpc.ServiceDesc for ForumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ForumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ForumService",
	HandlerType: (*ForumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateForum",
			Handler:    _ForumService_CreateForum_Handler,
		},
		{
			MethodName: "CreateThread",
			Handler:    _ForumService_CreateThread_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _ForumService_CreatePost_Handler,
		},
		{
			MethodName: "GetForum",
			Handler:    _ForumService_GetForum_Handler,
		},
		{
			MethodName: "GetThread",
			Handler:    _ForumService_GetThread_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _ForumService_GetPost_Handler,
		},
		{
			MethodName: "UpdateForum",
			Handler:    _ForumService_UpdateForum_Handler,
		},
		{
			MethodName: "UpdateThread",
			Handler:    _ForumService_UpdateThread_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _ForumService_UpdatePost_Handler,
		},
		{
			MethodName: "DeleteForum",
			Handler:    _ForumService_DeleteForum_Handler,
		},
		{
			MethodName: "DeleteThread",
			Handler:    _ForumService_DeleteThread_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _ForumService_DeletePost_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetForums",
			Handler:       _ForumService_GetForums_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetThreads",
			Handler:       _ForumService_GetThreads_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPosts",
			Handler:       _ForumService_GetPosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchForum",
			Handler:       _ForumService_SearchForum_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "forum.proto",
}
