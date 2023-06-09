// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: project.proto

package pbs

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

// HandleServerClient is the client API for HandleServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandleServerClient interface {
	ReviewProjectList(ctx context.Context, in *ReviewProjectListParams, opts ...grpc.CallOption) (*ReviewProjectListResponse, error)
	StreamClientServer(ctx context.Context, opts ...grpc.CallOption) (HandleServer_StreamClientServerClient, error)
}

type handleServerClient struct {
	cc grpc.ClientConnInterface
}

func NewHandleServerClient(cc grpc.ClientConnInterface) HandleServerClient {
	return &handleServerClient{cc}
}

func (c *handleServerClient) ReviewProjectList(ctx context.Context, in *ReviewProjectListParams, opts ...grpc.CallOption) (*ReviewProjectListResponse, error) {
	out := new(ReviewProjectListResponse)
	err := c.cc.Invoke(ctx, "/proto.HandleServer/ReviewProjectList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handleServerClient) StreamClientServer(ctx context.Context, opts ...grpc.CallOption) (HandleServer_StreamClientServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &HandleServer_ServiceDesc.Streams[0], "/proto.HandleServer/StreamClientServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &handleServerStreamClientServerClient{stream}
	return x, nil
}

type HandleServer_StreamClientServerClient interface {
	Send(*ParamId) error
	CloseAndRecv() (*ParamResp, error)
	grpc.ClientStream
}

type handleServerStreamClientServerClient struct {
	grpc.ClientStream
}

func (x *handleServerStreamClientServerClient) Send(m *ParamId) error {
	return x.ClientStream.SendMsg(m)
}

func (x *handleServerStreamClientServerClient) CloseAndRecv() (*ParamResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ParamResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HandleServerServer is the server API for HandleServer service.
// All implementations should embed UnimplementedHandleServerServer
// for forward compatibility
type HandleServerServer interface {
	ReviewProjectList(context.Context, *ReviewProjectListParams) (*ReviewProjectListResponse, error)
	StreamClientServer(HandleServer_StreamClientServerServer) error
}

// UnimplementedHandleServerServer should be embedded to have forward compatible implementations.
type UnimplementedHandleServerServer struct {
}

func (UnimplementedHandleServerServer) ReviewProjectList(context.Context, *ReviewProjectListParams) (*ReviewProjectListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewProjectList not implemented")
}
func (UnimplementedHandleServerServer) StreamClientServer(HandleServer_StreamClientServerServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClientServer not implemented")
}

// UnsafeHandleServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandleServerServer will
// result in compilation errors.
type UnsafeHandleServerServer interface {
	mustEmbedUnimplementedHandleServerServer()
}

func RegisterHandleServerServer(s grpc.ServiceRegistrar, srv HandleServerServer) {
	s.RegisterService(&HandleServer_ServiceDesc, srv)
}

func _HandleServer_ReviewProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewProjectListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleServerServer).ReviewProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HandleServer/ReviewProjectList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleServerServer).ReviewProjectList(ctx, req.(*ReviewProjectListParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandleServer_StreamClientServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HandleServerServer).StreamClientServer(&handleServerStreamClientServerServer{stream})
}

type HandleServer_StreamClientServerServer interface {
	SendAndClose(*ParamResp) error
	Recv() (*ParamId, error)
	grpc.ServerStream
}

type handleServerStreamClientServerServer struct {
	grpc.ServerStream
}

func (x *handleServerStreamClientServerServer) SendAndClose(m *ParamResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *handleServerStreamClientServerServer) Recv() (*ParamId, error) {
	m := new(ParamId)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HandleServer_ServiceDesc is the grpc.ServiceDesc for HandleServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandleServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HandleServer",
	HandlerType: (*HandleServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReviewProjectList",
			Handler:    _HandleServer_ReviewProjectList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClientServer",
			Handler:       _HandleServer_StreamClientServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "project.proto",
}
