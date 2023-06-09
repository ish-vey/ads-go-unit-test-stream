// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: unit_test_stream/unit-test-stream.proto

package unit_test_stream

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
	UnitTestStream_TestStream_FullMethodName = "/unit_test_stream.UnitTestStream/TestStream"
)

// UnitTestStreamClient is the client API for UnitTestStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnitTestStreamClient interface {
	// Sends bytes to server
	TestStream(ctx context.Context, opts ...grpc.CallOption) (UnitTestStream_TestStreamClient, error)
}

type unitTestStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewUnitTestStreamClient(cc grpc.ClientConnInterface) UnitTestStreamClient {
	return &unitTestStreamClient{cc}
}

func (c *unitTestStreamClient) TestStream(ctx context.Context, opts ...grpc.CallOption) (UnitTestStream_TestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UnitTestStream_ServiceDesc.Streams[0], UnitTestStream_TestStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &unitTestStreamTestStreamClient{stream}
	return x, nil
}

type UnitTestStream_TestStreamClient interface {
	Send(*VideoStreamRequest) error
	CloseAndRecv() (*VideoStreamReply, error)
	grpc.ClientStream
}

type unitTestStreamTestStreamClient struct {
	grpc.ClientStream
}

func (x *unitTestStreamTestStreamClient) Send(m *VideoStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *unitTestStreamTestStreamClient) CloseAndRecv() (*VideoStreamReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(VideoStreamReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UnitTestStreamServer is the server API for UnitTestStream service.
// All implementations must embed UnimplementedUnitTestStreamServer
// for forward compatibility
type UnitTestStreamServer interface {
	// Sends bytes to server
	TestStream(UnitTestStream_TestStreamServer) error
	mustEmbedUnimplementedUnitTestStreamServer()
}

// UnimplementedUnitTestStreamServer must be embedded to have forward compatible implementations.
type UnimplementedUnitTestStreamServer struct {
}

func (UnimplementedUnitTestStreamServer) TestStream(UnitTestStream_TestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TestStream not implemented")
}
func (UnimplementedUnitTestStreamServer) mustEmbedUnimplementedUnitTestStreamServer() {}

// UnsafeUnitTestStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnitTestStreamServer will
// result in compilation errors.
type UnsafeUnitTestStreamServer interface {
	mustEmbedUnimplementedUnitTestStreamServer()
}

func RegisterUnitTestStreamServer(s grpc.ServiceRegistrar, srv UnitTestStreamServer) {
	s.RegisterService(&UnitTestStream_ServiceDesc, srv)
}

func _UnitTestStream_TestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UnitTestStreamServer).TestStream(&unitTestStreamTestStreamServer{stream})
}

type UnitTestStream_TestStreamServer interface {
	SendAndClose(*VideoStreamReply) error
	Recv() (*VideoStreamRequest, error)
	grpc.ServerStream
}

type unitTestStreamTestStreamServer struct {
	grpc.ServerStream
}

func (x *unitTestStreamTestStreamServer) SendAndClose(m *VideoStreamReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *unitTestStreamTestStreamServer) Recv() (*VideoStreamRequest, error) {
	m := new(VideoStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UnitTestStream_ServiceDesc is the grpc.ServiceDesc for UnitTestStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UnitTestStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "unit_test_stream.UnitTestStream",
	HandlerType: (*UnitTestStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestStream",
			Handler:       _UnitTestStream_TestStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "unit_test_stream/unit-test-stream.proto",
}
