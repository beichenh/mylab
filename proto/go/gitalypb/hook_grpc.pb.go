// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gitalypb

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

// HookServiceClient is the client API for HookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HookServiceClient interface {
	PreReceiveHook(ctx context.Context, opts ...grpc.CallOption) (HookService_PreReceiveHookClient, error)
	PostReceiveHook(ctx context.Context, opts ...grpc.CallOption) (HookService_PostReceiveHookClient, error)
	UpdateHook(ctx context.Context, in *UpdateHookRequest, opts ...grpc.CallOption) (HookService_UpdateHookClient, error)
	ReferenceTransactionHook(ctx context.Context, opts ...grpc.CallOption) (HookService_ReferenceTransactionHookClient, error)
	// Deprecated: Do not use.
	// PackObjectsHook has been replaced by PackObjectsHookWithSidechannel. Remove in 15.0.
	PackObjectsHook(ctx context.Context, opts ...grpc.CallOption) (HookService_PackObjectsHookClient, error)
	// PackObjectsHookWithSidechannel is an optimized version of PackObjectsHook that uses
	// a unix socket side channel.
	PackObjectsHookWithSidechannel(ctx context.Context, in *PackObjectsHookWithSidechannelRequest, opts ...grpc.CallOption) (*PackObjectsHookWithSidechannelResponse, error)
}

type hookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHookServiceClient(cc grpc.ClientConnInterface) HookServiceClient {
	return &hookServiceClient{cc}
}

func (c *hookServiceClient) PreReceiveHook(ctx context.Context, opts ...grpc.CallOption) (HookService_PreReceiveHookClient, error) {
	stream, err := c.cc.NewStream(ctx, &HookService_ServiceDesc.Streams[0], "/gitaly.HookService/PreReceiveHook", opts...)
	if err != nil {
		return nil, err
	}
	x := &hookServicePreReceiveHookClient{stream}
	return x, nil
}

type HookService_PreReceiveHookClient interface {
	Send(*PreReceiveHookRequest) error
	Recv() (*PreReceiveHookResponse, error)
	grpc.ClientStream
}

type hookServicePreReceiveHookClient struct {
	grpc.ClientStream
}

func (x *hookServicePreReceiveHookClient) Send(m *PreReceiveHookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hookServicePreReceiveHookClient) Recv() (*PreReceiveHookResponse, error) {
	m := new(PreReceiveHookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hookServiceClient) PostReceiveHook(ctx context.Context, opts ...grpc.CallOption) (HookService_PostReceiveHookClient, error) {
	stream, err := c.cc.NewStream(ctx, &HookService_ServiceDesc.Streams[1], "/gitaly.HookService/PostReceiveHook", opts...)
	if err != nil {
		return nil, err
	}
	x := &hookServicePostReceiveHookClient{stream}
	return x, nil
}

type HookService_PostReceiveHookClient interface {
	Send(*PostReceiveHookRequest) error
	Recv() (*PostReceiveHookResponse, error)
	grpc.ClientStream
}

type hookServicePostReceiveHookClient struct {
	grpc.ClientStream
}

func (x *hookServicePostReceiveHookClient) Send(m *PostReceiveHookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hookServicePostReceiveHookClient) Recv() (*PostReceiveHookResponse, error) {
	m := new(PostReceiveHookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hookServiceClient) UpdateHook(ctx context.Context, in *UpdateHookRequest, opts ...grpc.CallOption) (HookService_UpdateHookClient, error) {
	stream, err := c.cc.NewStream(ctx, &HookService_ServiceDesc.Streams[2], "/gitaly.HookService/UpdateHook", opts...)
	if err != nil {
		return nil, err
	}
	x := &hookServiceUpdateHookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HookService_UpdateHookClient interface {
	Recv() (*UpdateHookResponse, error)
	grpc.ClientStream
}

type hookServiceUpdateHookClient struct {
	grpc.ClientStream
}

func (x *hookServiceUpdateHookClient) Recv() (*UpdateHookResponse, error) {
	m := new(UpdateHookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hookServiceClient) ReferenceTransactionHook(ctx context.Context, opts ...grpc.CallOption) (HookService_ReferenceTransactionHookClient, error) {
	stream, err := c.cc.NewStream(ctx, &HookService_ServiceDesc.Streams[3], "/gitaly.HookService/ReferenceTransactionHook", opts...)
	if err != nil {
		return nil, err
	}
	x := &hookServiceReferenceTransactionHookClient{stream}
	return x, nil
}

type HookService_ReferenceTransactionHookClient interface {
	Send(*ReferenceTransactionHookRequest) error
	Recv() (*ReferenceTransactionHookResponse, error)
	grpc.ClientStream
}

type hookServiceReferenceTransactionHookClient struct {
	grpc.ClientStream
}

func (x *hookServiceReferenceTransactionHookClient) Send(m *ReferenceTransactionHookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hookServiceReferenceTransactionHookClient) Recv() (*ReferenceTransactionHookResponse, error) {
	m := new(ReferenceTransactionHookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Deprecated: Do not use.
func (c *hookServiceClient) PackObjectsHook(ctx context.Context, opts ...grpc.CallOption) (HookService_PackObjectsHookClient, error) {
	stream, err := c.cc.NewStream(ctx, &HookService_ServiceDesc.Streams[4], "/gitaly.HookService/PackObjectsHook", opts...)
	if err != nil {
		return nil, err
	}
	x := &hookServicePackObjectsHookClient{stream}
	return x, nil
}

type HookService_PackObjectsHookClient interface {
	Send(*PackObjectsHookRequest) error
	Recv() (*PackObjectsHookResponse, error)
	grpc.ClientStream
}

type hookServicePackObjectsHookClient struct {
	grpc.ClientStream
}

func (x *hookServicePackObjectsHookClient) Send(m *PackObjectsHookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hookServicePackObjectsHookClient) Recv() (*PackObjectsHookResponse, error) {
	m := new(PackObjectsHookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hookServiceClient) PackObjectsHookWithSidechannel(ctx context.Context, in *PackObjectsHookWithSidechannelRequest, opts ...grpc.CallOption) (*PackObjectsHookWithSidechannelResponse, error) {
	out := new(PackObjectsHookWithSidechannelResponse)
	err := c.cc.Invoke(ctx, "/gitaly.HookService/PackObjectsHookWithSidechannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HookServiceServer is the server API for HookService service.
// All implementations must embed UnimplementedHookServiceServer
// for forward compatibility
type HookServiceServer interface {
	PreReceiveHook(HookService_PreReceiveHookServer) error
	PostReceiveHook(HookService_PostReceiveHookServer) error
	UpdateHook(*UpdateHookRequest, HookService_UpdateHookServer) error
	ReferenceTransactionHook(HookService_ReferenceTransactionHookServer) error
	// Deprecated: Do not use.
	// PackObjectsHook has been replaced by PackObjectsHookWithSidechannel. Remove in 15.0.
	PackObjectsHook(HookService_PackObjectsHookServer) error
	// PackObjectsHookWithSidechannel is an optimized version of PackObjectsHook that uses
	// a unix socket side channel.
	PackObjectsHookWithSidechannel(context.Context, *PackObjectsHookWithSidechannelRequest) (*PackObjectsHookWithSidechannelResponse, error)
	mustEmbedUnimplementedHookServiceServer()
}

// UnimplementedHookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHookServiceServer struct {
}

func (UnimplementedHookServiceServer) PreReceiveHook(HookService_PreReceiveHookServer) error {
	return status.Errorf(codes.Unimplemented, "method PreReceiveHook not implemented")
}
func (UnimplementedHookServiceServer) PostReceiveHook(HookService_PostReceiveHookServer) error {
	return status.Errorf(codes.Unimplemented, "method PostReceiveHook not implemented")
}
func (UnimplementedHookServiceServer) UpdateHook(*UpdateHookRequest, HookService_UpdateHookServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateHook not implemented")
}
func (UnimplementedHookServiceServer) ReferenceTransactionHook(HookService_ReferenceTransactionHookServer) error {
	return status.Errorf(codes.Unimplemented, "method ReferenceTransactionHook not implemented")
}
func (UnimplementedHookServiceServer) PackObjectsHook(HookService_PackObjectsHookServer) error {
	return status.Errorf(codes.Unimplemented, "method PackObjectsHook not implemented")
}
func (UnimplementedHookServiceServer) PackObjectsHookWithSidechannel(context.Context, *PackObjectsHookWithSidechannelRequest) (*PackObjectsHookWithSidechannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PackObjectsHookWithSidechannel not implemented")
}
func (UnimplementedHookServiceServer) mustEmbedUnimplementedHookServiceServer() {}

// UnsafeHookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HookServiceServer will
// result in compilation errors.
type UnsafeHookServiceServer interface {
	mustEmbedUnimplementedHookServiceServer()
}

func RegisterHookServiceServer(s grpc.ServiceRegistrar, srv HookServiceServer) {
	s.RegisterService(&HookService_ServiceDesc, srv)
}

func _HookService_PreReceiveHook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HookServiceServer).PreReceiveHook(&hookServicePreReceiveHookServer{stream})
}

type HookService_PreReceiveHookServer interface {
	Send(*PreReceiveHookResponse) error
	Recv() (*PreReceiveHookRequest, error)
	grpc.ServerStream
}

type hookServicePreReceiveHookServer struct {
	grpc.ServerStream
}

func (x *hookServicePreReceiveHookServer) Send(m *PreReceiveHookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hookServicePreReceiveHookServer) Recv() (*PreReceiveHookRequest, error) {
	m := new(PreReceiveHookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HookService_PostReceiveHook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HookServiceServer).PostReceiveHook(&hookServicePostReceiveHookServer{stream})
}

type HookService_PostReceiveHookServer interface {
	Send(*PostReceiveHookResponse) error
	Recv() (*PostReceiveHookRequest, error)
	grpc.ServerStream
}

type hookServicePostReceiveHookServer struct {
	grpc.ServerStream
}

func (x *hookServicePostReceiveHookServer) Send(m *PostReceiveHookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hookServicePostReceiveHookServer) Recv() (*PostReceiveHookRequest, error) {
	m := new(PostReceiveHookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HookService_UpdateHook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateHookRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HookServiceServer).UpdateHook(m, &hookServiceUpdateHookServer{stream})
}

type HookService_UpdateHookServer interface {
	Send(*UpdateHookResponse) error
	grpc.ServerStream
}

type hookServiceUpdateHookServer struct {
	grpc.ServerStream
}

func (x *hookServiceUpdateHookServer) Send(m *UpdateHookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HookService_ReferenceTransactionHook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HookServiceServer).ReferenceTransactionHook(&hookServiceReferenceTransactionHookServer{stream})
}

type HookService_ReferenceTransactionHookServer interface {
	Send(*ReferenceTransactionHookResponse) error
	Recv() (*ReferenceTransactionHookRequest, error)
	grpc.ServerStream
}

type hookServiceReferenceTransactionHookServer struct {
	grpc.ServerStream
}

func (x *hookServiceReferenceTransactionHookServer) Send(m *ReferenceTransactionHookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hookServiceReferenceTransactionHookServer) Recv() (*ReferenceTransactionHookRequest, error) {
	m := new(ReferenceTransactionHookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HookService_PackObjectsHook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HookServiceServer).PackObjectsHook(&hookServicePackObjectsHookServer{stream})
}

type HookService_PackObjectsHookServer interface {
	Send(*PackObjectsHookResponse) error
	Recv() (*PackObjectsHookRequest, error)
	grpc.ServerStream
}

type hookServicePackObjectsHookServer struct {
	grpc.ServerStream
}

func (x *hookServicePackObjectsHookServer) Send(m *PackObjectsHookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hookServicePackObjectsHookServer) Recv() (*PackObjectsHookRequest, error) {
	m := new(PackObjectsHookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HookService_PackObjectsHookWithSidechannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackObjectsHookWithSidechannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).PackObjectsHookWithSidechannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.HookService/PackObjectsHookWithSidechannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).PackObjectsHookWithSidechannel(ctx, req.(*PackObjectsHookWithSidechannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HookService_ServiceDesc is the grpc.ServiceDesc for HookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.HookService",
	HandlerType: (*HookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PackObjectsHookWithSidechannel",
			Handler:    _HookService_PackObjectsHookWithSidechannel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PreReceiveHook",
			Handler:       _HookService_PreReceiveHook_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PostReceiveHook",
			Handler:       _HookService_PostReceiveHook_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateHook",
			Handler:       _HookService_UpdateHook_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReferenceTransactionHook",
			Handler:       _HookService_ReferenceTransactionHook_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PackObjectsHook",
			Handler:       _HookService_PackObjectsHook_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hook.proto",
}
