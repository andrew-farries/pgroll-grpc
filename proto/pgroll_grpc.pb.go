// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/pgroll.proto

package proto

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

const (
	PGRoll_Start_FullMethodName    = "/pgroll.PGRoll/Start"
	PGRoll_Complete_FullMethodName = "/pgroll.PGRoll/Complete"
	PGRoll_Rollback_FullMethodName = "/pgroll.PGRoll/Rollback"
)

// PGRollClient is the client API for PGRoll service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PGRollClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (PGRoll_StartClient, error)
	Complete(ctx context.Context, in *CompleteRequest, opts ...grpc.CallOption) (*CompleteResponse, error)
	Rollback(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*RollbackResponse, error)
}

type pGRollClient struct {
	cc grpc.ClientConnInterface
}

func NewPGRollClient(cc grpc.ClientConnInterface) PGRollClient {
	return &pGRollClient{cc}
}

func (c *pGRollClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (PGRoll_StartClient, error) {
	stream, err := c.cc.NewStream(ctx, &PGRoll_ServiceDesc.Streams[0], PGRoll_Start_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pGRollStartClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PGRoll_StartClient interface {
	Recv() (*StartResponse, error)
	grpc.ClientStream
}

type pGRollStartClient struct {
	grpc.ClientStream
}

func (x *pGRollStartClient) Recv() (*StartResponse, error) {
	m := new(StartResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pGRollClient) Complete(ctx context.Context, in *CompleteRequest, opts ...grpc.CallOption) (*CompleteResponse, error) {
	out := new(CompleteResponse)
	err := c.cc.Invoke(ctx, PGRoll_Complete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pGRollClient) Rollback(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*RollbackResponse, error) {
	out := new(RollbackResponse)
	err := c.cc.Invoke(ctx, PGRoll_Rollback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PGRollServer is the server API for PGRoll service.
// All implementations must embed UnimplementedPGRollServer
// for forward compatibility
type PGRollServer interface {
	Start(*StartRequest, PGRoll_StartServer) error
	Complete(context.Context, *CompleteRequest) (*CompleteResponse, error)
	Rollback(context.Context, *RollbackRequest) (*RollbackResponse, error)
	mustEmbedUnimplementedPGRollServer()
}

// UnimplementedPGRollServer must be embedded to have forward compatible implementations.
type UnimplementedPGRollServer struct {
}

func (UnimplementedPGRollServer) Start(*StartRequest, PGRoll_StartServer) error {
	return status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedPGRollServer) Complete(context.Context, *CompleteRequest) (*CompleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complete not implemented")
}
func (UnimplementedPGRollServer) Rollback(context.Context, *RollbackRequest) (*RollbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rollback not implemented")
}
func (UnimplementedPGRollServer) mustEmbedUnimplementedPGRollServer() {}

// UnsafePGRollServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PGRollServer will
// result in compilation errors.
type UnsafePGRollServer interface {
	mustEmbedUnimplementedPGRollServer()
}

func RegisterPGRollServer(s grpc.ServiceRegistrar, srv PGRollServer) {
	s.RegisterService(&PGRoll_ServiceDesc, srv)
}

func _PGRoll_Start_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StartRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PGRollServer).Start(m, &pGRollStartServer{stream})
}

type PGRoll_StartServer interface {
	Send(*StartResponse) error
	grpc.ServerStream
}

type pGRollStartServer struct {
	grpc.ServerStream
}

func (x *pGRollStartServer) Send(m *StartResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PGRoll_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PGRollServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PGRoll_Complete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PGRollServer).Complete(ctx, req.(*CompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PGRoll_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PGRollServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PGRoll_Rollback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PGRollServer).Rollback(ctx, req.(*RollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PGRoll_ServiceDesc is the grpc.ServiceDesc for PGRoll service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PGRoll_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pgroll.PGRoll",
	HandlerType: (*PGRollServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Complete",
			Handler:    _PGRoll_Complete_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _PGRoll_Rollback_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Start",
			Handler:       _PGRoll_Start_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/pgroll.proto",
}
